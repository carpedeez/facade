package facade

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Get Session
// (GET /@me)
func (f facade) Me(w http.ResponseWriter, r *http.Request) *Response {
	cookies := r.Header.Get("Cookie")

	s, _, err := f.ory.V0alpha2Api.ToSession(r.Context()).Cookie(cookies).Execute()
	if (err != nil || s == nil) || !*s.Active {
		return ErrorResponse("Unauthorized", http.StatusUnauthorized)
	}
	b, err := json.Marshal(s)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	ss := Session{}
	json.Unmarshal(b, &ss)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	return MeJSON200Response(ss)
}

// Upload file
// (POST /assets)
func (f facade) UploadFile(w http.ResponseWriter, r *http.Request) *Response {
	return UploadFileJSON200Response("google.com")
}

// Create display
// (POST /d)
func (f facade) CreateDisplay(w http.ResponseWriter, r *http.Request) *Response {
	s := getSession(r.Context())
	display := PostDisplay{}
	err := json.NewDecoder(r.Body).Decode(&display)
	if err != nil {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if display.Description == "" || display.Title == "" {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	id, err := f.querier.CreateDisplay(s.Identity.Id, display.Title, display.Description) // should we respond with the serial number and redirect them? should we respond with the whole object?
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	return CreateDisplayJSON200Response(id)
}

// Delete display
// (DELETE /d/{displayID})
func (f facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	s := getSession(r.Context())
	isDisplayOwner, err := f.querier.IsDisplayOwner(s.Identity.Id, displayID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isDisplayOwner {
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}
	err = f.querier.DeleteDisplay(displayID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	return NoContentResponse()
}

// Get display
// (GET /d/{displayID})
func (f facade) GetDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	d, err := f.querier.GetDisplay(displayID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	items, err := f.querier.GetItems(displayID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	displayItems := make([]GetItem, len(items))
	for _, i := range items {
		displayItems = append(displayItems, GetItem{
			DisplayID:      displayID,
			ExternalLink:   i.ExternalLink,
			ID:             i.ID,
			PhotoURL:       i.PhotoURL,
			SocialPostLink: i.SocialPostLink,
			UserID:         d.UserID,
		})
	}
	return GetDisplayJSON200Response(GetDisplay{
		Description: d.Description,
		ID:          displayID,
		Items:       displayItems,
		PhotoURL:    d.PhotoURL,
		Title:       d.Title,
		UserID:      d.UserID,
	})
}

// Update display
// (PATCH /d/{displayID})
func (f facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	s := getSession(r.Context())
	isDisplayOwner, err := f.querier.IsDisplayOwner(s.Identity.Id, displayID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isDisplayOwner {
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}
	display := PatchDisplay{}
	err = json.NewDecoder(r.Body).Decode(&display)
	if err != nil {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if display.Description == nil && display.Title == nil {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	updatedDisplay, err := f.querier.UpdateDisplay(displayID, display.Title, display.Description) // should we respond with the serial number and redirect them? should we respond with the whole object?
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}

	items, err := f.querier.GetItems(displayID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	displayItems := make([]GetItem, len(items))
	for _, i := range items {
		displayItems = append(displayItems, GetItem{
			DisplayID:      displayID,
			ExternalLink:   i.ExternalLink,
			ID:             i.ID,
			PhotoURL:       i.PhotoURL,
			SocialPostLink: i.SocialPostLink,
			UserID:         s.Identity.Id,
		})
	}

	return GetDisplayJSON200Response(GetDisplay{
		Description: updatedDisplay.Description,
		ID:          displayID,
		Items:       displayItems,
		PhotoURL:    updatedDisplay.PhotoURL,
		Title:       updatedDisplay.Title,
		UserID:      updatedDisplay.UserID,
	})
}

// Create Item
// (POST /i)
func (f facade) CreateItem(w http.ResponseWriter, r *http.Request) *Response {
	s := getSession(r.Context())
	item := PostItem{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if item.ExternalLink == "" {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}

	isDisplayOwner, err := f.querier.IsDisplayOwner(s.Identity.Id, item.DisplayID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isDisplayOwner {
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}
	createdItem, err := f.querier.CreateItem(s.Identity.Id, item.DisplayID, item.ExternalLink) // should we respond with the serial number and redirect them? should we respond with the whole object?
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	return CreateItemJSON200Response(GetItem{
		DisplayID:      item.DisplayID,
		ExternalLink:   item.ExternalLink,
		ID:             createdItem.ID,
		PhotoURL:       createdItem.PhotoURL,
		SocialPostLink: createdItem.SocialPostLink,
		UserID:         createdItem.UserID,
	})
}

// Delete item
// (DELETE /i/{itemID})
func (f facade) DeleteItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	s := getSession(r.Context())
	isItemOwner, err := f.querier.IsItemOwner(s.Identity.Id, itemID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isItemOwner {
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}
	err = f.querier.DeleteItem(itemID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	return NoContentResponse()
}

// Get item
// (GET /i/{itemID})
func (f facade) GetItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	i, err := f.querier.GetItem(itemID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrorResponse("Not Found", http.StatusNotFound)
		}
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	return GetItemJSON200Response(GetItem{
		DisplayID:      int64(i.DisplayID),
		ExternalLink:   i.ExternalLink,
		ID:             itemID,
		PhotoURL:       i.PhotoURL,
		SocialPostLink: i.SocialPostLink,
		UserID:         i.UserID,
	})
}

// Update item
// (PATCH /i/{itemID})
func (f facade) UpdateItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	s := getSession(r.Context())
	isItemOwner, err := f.querier.IsItemOwner(s.Identity.Id, itemID)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	if !isItemOwner {
		return ErrorResponse("Forbidden", http.StatusForbidden)
	}
	item := PatchItem{}
	err = json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	if item.ExternalLink == nil && item.PhotoURL == nil && item.SocialPostLink == nil {
		return ErrorResponse("Bad Request", http.StatusBadRequest)
	}
	updatedItem, err := f.querier.UpdateItem(itemID, item.ExternalLink, item.SocialPostLink, item.PhotoURL)
	if err != nil {
		return ErrorResponse("Internal Server Error", http.StatusInternalServerError)
	}
	return UpdateItemJSON200Response(GetItem{
		DisplayID:      updatedItem.DisplayID,
		ExternalLink:   updatedItem.ExternalLink,
		ID:             itemID,
		PhotoURL:       updatedItem.PhotoURL,
		SocialPostLink: updatedItem.SocialPostLink,
		UserID:         updatedItem.UserID,
	})
}

func ErrorResponse(message string, code int) *Response {
	b, err := json.Marshal(Error{Message: message, Code: int32(code)})
	if err != nil {
		return &Response{
			Code: http.StatusInternalServerError,
			body: "Internal Server Error",
		}
	}
	return &Response{
		Code: code,
		body: b,
	}
}

func NoContentResponse() *Response {
	return &Response{
		Code: http.StatusNoContent,
	}
}
