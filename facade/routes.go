package facade

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Get Session
// (GET /@me)
func (f facade) Me(w http.ResponseWriter, r *http.Request) *Response {
	cookies := r.Header.Get("Cookie")
	s, _, err := f.ory.V0alpha2Api.ToSession(r.Context()).Cookie(cookies).Execute()
	if (err != nil || s == nil) || !*s.Active {
		return ErrorResponse("Unauthorized", http.StatusUnauthorized)
	}
	return &Response{ // switched to this instead of the MeJSON200Response because it's wrong.
		body: s, // fix open api schema when identity schema is decided
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

	d, err := f.querier.CreateDisplay(s.Identity.Id, pd.Title, pd.Description) // should we respond with the serial number and redirect them? should we respond with the whole object?
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to create display in database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Add("Location", "/d/"+strconv.Itoa(int(d.ID)))
	return CreateDisplayJSON201Response(GetDisplay{
		Description: d.Description,
		ID:          d.ID,
		Items:       []GetItem{},
		PhotoURL:    d.PhotoURL,
		Title:       d.Title,
		UserID:      d.UserID,
	})
}

// Get display
// (GET /d/{displayID})
func (f facade) GetDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	d, err := f.querier.GetDisplay(displayID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get display from database: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	is, err := f.querier.GetItems(displayID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get items from database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	dItems := []GetItem{}
	for _, i := range is {
		dItems = append(dItems, GetItem{
			DisplayID:      i.DisplayID,
			ExternalLink:   i.ExternalLink,
			ID:             i.ID,
			PhotoURL:       i.PhotoURL,
			SocialPostLink: i.SocialPostLink,
			UserID:         i.UserID,
		})
	}

	return GetDisplayJSON200Response(GetDisplay{
		Description: d.Description,
		ID:          d.ID,
		Items:       dItems,
		PhotoURL:    d.PhotoURL,
		Title:       d.Title,
		UserID:      d.UserID,
	})
}

// Update display
// (PATCH /d/{displayID})
func (f facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	s := getSession(r.Context())

	pd := PatchDisplay{}
	err := json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to decode patchdisplay body into a patchdisplay: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if pd.Description == nil && pd.Title == nil {
		f.log.Error().Err(fmt.Errorf("user did not provide either a title or description: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	isDisplayOwner, err := f.querier.IsDisplayOwner(s.Identity.Id, displayID)
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

	d, err := f.querier.UpdateDisplay(displayID, pd.Title, pd.Description)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to update display in database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	is, err := f.querier.GetItems(displayID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get items from database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	dItems := []GetItem{}
	for _, i := range is {
		dItems = append(dItems, GetItem{
			DisplayID:      i.DisplayID,
			ExternalLink:   i.ExternalLink,
			ID:             i.ID,
			PhotoURL:       i.PhotoURL,
			SocialPostLink: i.SocialPostLink,
			UserID:         i.UserID,
		})
	}

	return GetDisplayJSON200Response(GetDisplay{
		Description: d.Description,
		ID:          d.ID,
		Items:       dItems,
		PhotoURL:    d.PhotoURL,
		Title:       d.Title,
		UserID:      d.UserID,
	})
}

// Delete display
// (DELETE /d/{displayID})
func (f facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	s := getSession(r.Context())

	isDisplayOwner, err := f.querier.IsDisplayOwner(s.Identity.Id, displayID)
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

	ok, err := f.querier.DeleteDisplay(displayID)
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
func (f facade) CreateItem(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	s := getSession(r.Context())

	pi := PostItem{}
	err := json.NewDecoder(r.Body).Decode(&pi)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to decode postitem body into a postitem: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if pi.ExternalLink == "" {
		f.log.Error().Err(fmt.Errorf("user did not provide an external link: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	isDisplayOwner, err := f.querier.IsDisplayOwner(s.Identity.Id, displayID)
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

	i, err := f.querier.CreateItem(s.Identity.Id, displayID, pi.ExternalLink) // should we respond with the serial number and redirect them? should we respond with the whole object?
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to create item in database: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	w.Header().Add("Location", "/i/"+strconv.Itoa(int(i.ID)))
	return CreateItemJSON201Response(GetItem{
		DisplayID:      i.DisplayID,
		ExternalLink:   i.ExternalLink,
		ID:             i.ID,
		PhotoURL:       i.PhotoURL,
		SocialPostLink: i.SocialPostLink,
		UserID:         i.UserID,
	})
}

// Get item
// (GET /d/{displayID}/i/{itemID})
func (f facade) GetItem(w http.ResponseWriter, r *http.Request, displayID int64, itemID int64) *Response {
	i, err := f.querier.GetItem(itemID, displayID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to get item from database: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	return GetItemJSON200Response(GetItem{
		DisplayID:      i.DisplayID,
		ExternalLink:   i.ExternalLink,
		ID:             i.ID,
		PhotoURL:       i.PhotoURL,
		SocialPostLink: i.SocialPostLink,
		UserID:         i.UserID,
	})
}

// Update item
// (PATCH /d/{displayID}/i/{itemID})
func (f facade) UpdateItem(w http.ResponseWriter, r *http.Request, displayID int64, itemID int64) *Response {
	s := getSession(r.Context())

	pi := PatchItem{}
	err := json.NewDecoder(r.Body).Decode(&pi)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to decode patchitem body into a patchitem: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if pi.ExternalLink == nil && pi.PhotoURL == nil && pi.SocialPostLink == nil {
		f.log.Error().Err(fmt.Errorf("user did not provide a externallink, photourl, or socialpostlink: %w", err)).Msg("")
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	isItemOwner, err := f.querier.IsItemOwner(s.Identity.Id, itemID)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to check if user is item owner: %w", err)).Msg("")
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isItemOwner {
		f.log.Error().Err(fmt.Errorf("user attempted to update an item they do not own: %w", err)).Msg("")
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}

	i, err := f.querier.UpdateItem(itemID, displayID, pi.ExternalLink, pi.SocialPostLink, pi.PhotoURL)
	if err != nil {
		f.log.Error().Err(fmt.Errorf("failed to update item in database: %w", err)).Msg("")
		if errors.Is(err, sql.ErrNoRows) {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	return UpdateItemJSON200Response(GetItem{
		DisplayID:      i.DisplayID,
		ExternalLink:   i.ExternalLink,
		ID:             i.ID,
		PhotoURL:       i.PhotoURL,
		SocialPostLink: i.SocialPostLink,
		UserID:         i.UserID,
	})
}

// Delete item
// (DELETE /d/{displayID}/i/{itemID})
func (f facade) DeleteItem(w http.ResponseWriter, r *http.Request, displayID int64, itemID int64) *Response {
	s := getSession(r.Context())

	isItemOwner, err := f.querier.IsItemOwner(s.Identity.Id, itemID)
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

	ok, err := f.querier.DeleteItem(itemID, displayID)
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
