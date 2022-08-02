package facade

import (
	"encoding/json"
	"net/http"

	"github.com/carpedeez/store/database"
)

type Facade struct {
	Querier *database.Querier // consider making Querier an interface, with querierImpl being the implementation
}

// Upload file
// (POST /assets)
func (f Facade) UploadFile(w http.ResponseWriter, r *http.Request) *Response {
	return UploadFileJSON200Response("google.com")
}

// Create display
// (POST /d)
func (f Facade) CreateDisplay(w http.ResponseWriter, r *http.Request) *Response {
	display := Display{}
	err := json.NewDecoder(r.Body).Decode(&display)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	if display.Description == "" || display.Title == "" || display.Username == "" { // dereferencing nil pointers = panic, so we gotta check before we call CreateDisplay
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	id, err := f.Querier.CreateDisplay(display.Username, display.Title, display.Description) // should we respond with the serial number and redirect them? should we respond with the whole object?
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
func (f Facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Get display
// (GET /d/{displayID})
func (f Facade) GetDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	return GetDisplayJSON200Response(Display{
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
func (f Facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	return UpdateDisplayJSON200Response(Display{
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
func (f Facade) CreateItem(w http.ResponseWriter, r *http.Request) *Response {
	return CreateItemJSON200Response(Item{
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
func (f Facade) DeleteItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Get item
// (GET /i/{itemID})
func (f Facade) GetItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	return GetItemJSON200Response(Item{
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
func (f Facade) UpdateItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	return UpdateItemJSON200Response(Item{
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
func (f Facade) GetUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return GetUserJSON200Response(User{
		DisplayIDS:  []int64{},
		Fname:       "",
		Lname:       "",
		PhotoURL:    "",
		SocialLinks: []string{},
		Username:    "",
	})
}

// Update user
// (PATCH /{username})
func (f Facade) UpdateUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return UpdateUserJSON200Response(User{
		DisplayIDS:  []int64{},
		Fname:       "",
		Lname:       "",
		PhotoURL:    "",
		SocialLinks: []string{},
		Username:    "",
	})
}

// Create user
// (POST /{username})
func (f Facade) CreateUser(w http.ResponseWriter, r *http.Request) *Response {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	if user.Username == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	err = f.Querier.CreateUser(user.Username, user.Fname, user.Lname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
