package facade

import (
	"encoding/json"
	"net/http"
)

// Upload file
// (POST /assets)
func (f facade) UploadFile(w http.ResponseWriter, r *http.Request) *Response {
	return UploadFileJSON200Response("google.com")
}

// Create display
// (POST /d)
func (f facade) CreateDisplay(w http.ResponseWriter, r *http.Request) *Response {
	display := PostDisplay{}
	err := json.NewDecoder(r.Body).Decode(&display)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	if display.Description == "" || display.Title == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	id, err := f.querier.CreateDisplay(1, display.Title, display.Description) // should we respond with the serial number and redirect them? should we respond with the whole object?
	if err != nil {
		return CreateDisplayJSONDefaultResponse(Error{
			Code:    http.StatusBadRequest, // this is placeholder, we need to figure out what errors we should respond with. also, is the error object (as opposed to normal status codes) ideal?
			Message: err.Error(),
		})
	}

	return CreateDisplayJSON200Response(id)
}

// Delete display
// (DELETE /d/{displayID})
func (f facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Get display
// (GET /d/{displayID})
func (f facade) GetDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	return GetDisplayJSON200Response(GetDisplay{
		Description: "",
		ID:          0,
		ItemIDS:     []int64{},
		PhotoURL:    "",
		Title:       "",
		Username:    "",
	})
}

// Update display
// (PATCH /d/{displayID})
func (f facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	return UpdateDisplayJSON200Response(GetDisplay{
		Description: "",
		ID:          0,
		ItemIDS:     []int64{},
		PhotoURL:    "",
		Title:       "",
		Username:    "",
	})
}

// Create Item
// (POST /i)
func (f facade) CreateItem(w http.ResponseWriter, r *http.Request) *Response {
	return CreateItemJSON200Response(GetItem{
		DisplayID:      0,
		ExternalLink:   "",
		ID:             0,
		PhotoURL:       "",
		SocialPostLink: "",
		Username:       "",
	})
}

// Delete item
// (DELETE /i/{itemID})
func (f facade) DeleteItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Get item
// (GET /i/{itemID})
func (f facade) GetItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	return GetItemJSON200Response(GetItem{
		DisplayID:      0,
		ExternalLink:   "",
		ID:             0,
		PhotoURL:       "",
		SocialPostLink: "",
		Username:       "",
	})
}

// Update item
// (PATCH /i/{itemID})
func (f facade) UpdateItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	return UpdateItemJSON200Response(GetItem{
		DisplayID:      0,
		ExternalLink:   "",
		ID:             0,
		PhotoURL:       "",
		SocialPostLink: "",
		Username:       "",
	})
}

// Get user
// (GET /{username})
func (f facade) GetUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return GetUserJSON200Response(GetUser{
		DisplayIDS:  []int64{},
		FirstName:   "",
		LastName:    "",
		PhotoURL:    "",
		SocialLinks: []string{},
		Username:    "",
	})
}

// Update user
// (PATCH /{username})
func (f facade) UpdateUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return UpdateUserJSON200Response(GetUser{
		DisplayIDS:  []int64{},
		FirstName:   "",
		LastName:    "",
		PhotoURL:    "",
		SocialLinks: []string{},
		Username:    "",
	})
}

// Create user
// (POST /{username})
func (f facade) CreateUser(w http.ResponseWriter, r *http.Request) *Response {
	user := PostUser{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	if user.Username == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	err = f.querier.CreateUser(user.Username, user.FirstName, user.LastName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
