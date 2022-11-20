package facade

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// Get Session
// (GET /@me)
func (f facade) Me(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		body: r.Context().Value("req.session"), // fix open api schema when identity schema is decided
		Code: 200,
	}
}

// Upload file
// (POST /assets)
func (f facade) UploadFile(w http.ResponseWriter, r *http.Request) *Response {
	// we will probably want 201 with the Location header
	return UploadFileJSON200Response("google.com")
}

// Create display
// (POST /d)
func (f facade) CreateDisplay(w http.ResponseWriter, r *http.Request) *Response {
	s := getSession(r.Context())
	uID, _ := uuid.Parse(s.Identity.Id) // Prayge Kratos give us a real UUID

	pd := PostDisplay{}
	err := json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to decode postdisplay body into a postdisplay: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	if pd.Description == "" || pd.Title == "" {
		f.log.Error().Err(fmt.Errorf("user did not provide either a title, description, or either: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	d, err := f.querier.CreateDisplay(uID, pd.Title, pd.Description)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to create display in database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Add("Location", "/d/"+d.ID.String())
	return CreateDisplayJSON201Response(GetDisplay{
		ID:          d.ID.String(),
		UserID:      d.UserID.String(),
		Title:       d.Title,
		Description: d.Description,
		PhotoURL:    d.PhotoURL,
		Items:       []GetItem{},
	})
}

// Get display
// (GET /d/{displayID})
func (f facade) GetDisplay(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response {
	dID, err := uuid.Parse(string(displayID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	d, err := f.querier.GetDisplay(dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get display from database: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	is, err := f.querier.GetItems(dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get items from database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	dItems := []GetItem{}
	for _, i := range is {
		dItems = append(dItems, GetItem{
			ID:             i.ID.String(),
			DisplayID:      i.DisplayID.String(),
			UserID:         i.UserID.String(),
			ExternalLink:   i.ExternalLink,
			SocialPostLink: i.SocialPostLink,
			PhotoURL:       i.PhotoURL,
		})
	}

	return GetDisplayJSON200Response(GetDisplay{
		ID:          d.ID.String(),
		UserID:      d.UserID.String(),
		Title:       d.Title,
		Description: d.Description,
		PhotoURL:    d.PhotoURL,
		Items:       dItems,
	})
}

// Update display
// (PATCH /d/{displayID})
func (f facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response {
	s := getSession(r.Context())
	uID, _ := uuid.Parse(s.Identity.Id) // Prayge Kratos give us a real UUID

	dID, err := uuid.Parse(string(displayID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	pd := PatchDisplay{}
	err = json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to decode patchdisplay body into a patchdisplay: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if pd.Description == nil && pd.Title == nil {
		f.log.Error().Err(fmt.Errorf("user did not provide either a title or description: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	isDisplayOwner, err := f.querier.IsDisplayOwner(uID, dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to check if user is display owner: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isDisplayOwner {
		f.log.Error().Err(fmt.Errorf("user attempted to update a display they do not own: %w", err)).Msg("")
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}

	d, err := f.querier.UpdateDisplay(dID, pd.Title, pd.Description)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to update display in database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	is, err := f.querier.GetItems(dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get items from database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	dItems := []GetItem{}
	for _, i := range is {
		dItems = append(dItems, GetItem{
			ID:             i.ID.String(),
			UserID:         i.UserID.String(),
			DisplayID:      i.DisplayID.String(),
			ExternalLink:   i.ExternalLink,
			SocialPostLink: i.SocialPostLink,
			PhotoURL:       i.PhotoURL,
		})
	}

	return GetDisplayJSON200Response(GetDisplay{
		ID:          d.ID.String(),
		UserID:      d.UserID.String(),
		Title:       d.Title,
		Description: d.Description,
		PhotoURL:    d.PhotoURL,
		Items:       dItems,
	})
}

// Delete display
// (DELETE /d/{displayID})
func (f facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response {
	s := getSession(r.Context())
	uID, _ := uuid.Parse(s.Identity.Id) // Prayge Kratos give us a real UUID

	dID, err := uuid.Parse(string(displayID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	isDisplayOwner, err := f.querier.IsDisplayOwner(uID, dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to check if user is display owner: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isDisplayOwner {
		f.log.Error().Err(fmt.Errorf("user attempted to delete a display they do not own: %w", err)).Msg("")
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}

	ok, err := f.querier.DeleteDisplay(dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to delete display from database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	if !ok {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	return NoContentResponse()
}

// Create Item
// (POST /d/{displayID}/i)
func (f facade) CreateItem(w http.ResponseWriter, r *http.Request, displayID DisplayID) *Response {
	s := getSession(r.Context())
	uID, _ := uuid.Parse(s.Identity.Id) // Prayge Kratos give us a real UUID

	dID, err := uuid.Parse(string(displayID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	pi := PostItem{}
	err = json.NewDecoder(r.Body).Decode(&pi)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to decode postitem body into a postitem: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if pi.ExternalLink == "" {
		f.log.Error().Err(fmt.Errorf("user did not provide an external link: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	isDisplayOwner, err := f.querier.IsDisplayOwner(uID, dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to check if user is display owner: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isDisplayOwner {
		f.log.Error().Err(fmt.Errorf("user attempted to create an item in a display they do not own: %w", err)).Msg("")
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}

	i, err := f.querier.CreateItem(uID, dID, pi.ExternalLink)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to create item in database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Add("Location", "/i/"+i.ID.String())
	return CreateItemJSON201Response(GetItem{
		ID:             i.ID.String(),
		UserID:         i.UserID.String(),
		DisplayID:      i.DisplayID.String(),
		ExternalLink:   i.ExternalLink,
		SocialPostLink: i.SocialPostLink,
		PhotoURL:       i.PhotoURL,
	})
}

// Get item
// (GET /d/{displayID}/i/{itemID})
func (f facade) GetItem(w http.ResponseWriter, r *http.Request, displayID DisplayID, itemID ItemID) *Response {
	dID, err := uuid.Parse(string(displayID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}
	iID, err := uuid.Parse(string(itemID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	i, err := f.querier.GetItem(iID, dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get item from database: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	return GetItemJSON200Response(GetItem{
		ID:             i.ID.String(),
		UserID:         i.UserID.String(),
		DisplayID:      i.DisplayID.String(),
		ExternalLink:   i.ExternalLink,
		SocialPostLink: i.SocialPostLink,
		PhotoURL:       i.PhotoURL,
	})
}

// Update item
// (PATCH /d/{displayID}/i/{itemID})
func (f facade) UpdateItem(w http.ResponseWriter, r *http.Request, displayID DisplayID, itemID ItemID) *Response {
	s := getSession(r.Context())
	uID, _ := uuid.Parse(s.Identity.Id) // Prayge Kratos give us a real UUID

	dID, err := uuid.Parse(string(displayID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}
	iID, err := uuid.Parse(string(itemID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	pi := PatchItem{}
	err = json.NewDecoder(r.Body).Decode(&pi)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to decode patchitem body into a patchitem: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if pi.ExternalLink == nil && pi.PhotoURL == nil && pi.SocialPostLink == nil {
		f.log.Error().Err(fmt.Errorf("user did not provide a externallink, photourl, or socialpostlink: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	isItemOwner, err := f.querier.IsItemOwner(uID, iID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to check if user is item owner: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isItemOwner {
		f.log.Error().Err(fmt.Errorf("user attempted to update an item they do not own: %w", err)).Msg("")
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}

	i, err := f.querier.UpdateItem(iID, dID, pi.ExternalLink, pi.SocialPostLink, pi.PhotoURL)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to update item in database: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	return UpdateItemJSON200Response(GetItem{
		ID:             i.ID.String(),
		UserID:         i.UserID.String(),
		DisplayID:      i.DisplayID.String(),
		ExternalLink:   i.ExternalLink,
		SocialPostLink: i.SocialPostLink,
		PhotoURL:       i.PhotoURL,
	})
}

// Delete item
// (DELETE /d/{displayID}/i/{itemID})
func (f facade) DeleteItem(w http.ResponseWriter, r *http.Request, displayID DisplayID, itemID ItemID) *Response {
	s := getSession(r.Context())
	uID, _ := uuid.Parse(s.Identity.Id) // Prayge Kratos give us a real UUID

	dID, err := uuid.Parse(string(displayID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}
	iID, err := uuid.Parse(string(itemID))
	if err != nil {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	isItemOwner, err := f.querier.IsItemOwner(uID, iID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to check if user is item owner: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isItemOwner {
		f.log.Error().Err(fmt.Errorf("user attempted to delete an item they do not own: %w", err)).Msg("")
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}

	ok, err := f.querier.DeleteItem(iID, dID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to delete item from database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	if !ok {
		return ErrorResponse("Not Found", http.StatusNotFound)
	}

	return NoContentResponse()
}

func ErrorResponse(message string, code int) *Response {
	return &Response{
		Code: code,
		body: Error{Message: message, Code: int32(code)},
	}
}

func NoContentResponse() *Response {
	return &Response{
		Code: http.StatusNoContent,
	}
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
