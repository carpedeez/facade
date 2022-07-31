package facade

import (
	"net/http"

	"github.com/carpedeez/store/database"
)

type Facade struct {
	Querier *database.Querier
}

// Upload file
// (POST /assets)
func (f Facade) UploadFile(w http.ResponseWriter, r *http.Request) *Response {
	return UploadFileJSON200Response("google.com")
}

// Create display
// (POST /d)
func (f Facade) CreateDisplay(w http.ResponseWriter, r *http.Request, params CreateDisplayParams) *Response {
	return CreateDisplayJSON200Response(Display{
		Description: new(string),
		ID:          new(uint64),
		ItemIDS:     []uint64{},
		PhotoURL:    new(string),
		Title:       new(string),
		Username:    new(string),
	})
}

// Delete display
// (DELETE /d/{displayID})
func (f Facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) *Response {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Get display
// (GET /d/{displayID})
func (f Facade) GetDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) *Response {
	return GetDisplayJSON200Response(Display{
		Description: new(string),
		ID:          new(uint64),
		ItemIDS:     []uint64{},
		PhotoURL:    new(string),
		Title:       new(string),
		Username:    new(string),
	})
}

// Update display
// (PATCH /d/{displayID})
func (f Facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) *Response {
	return UpdateDisplayJSON200Response(Display{
		Description: new(string),
		ID:          new(uint64),
		ItemIDS:     []uint64{},
		PhotoURL:    new(string),
		Title:       new(string),
		Username:    new(string),
	})
}

// Create Item
// (POST /i)
func (f Facade) CreateItem(w http.ResponseWriter, r *http.Request, params CreateItemParams) *Response {
	return CreateItemJSON200Response(Item{
		DisplayID:      new(uint64),
		ExternalLink:   new(string),
		ID:             new(uint64),
		PhotoURL:       new(string),
		SocialPostLink: new(string),
		Username:       new(string),
	})
}

// Delete item
// (DELETE /i/{itemID})
func (f Facade) DeleteItem(w http.ResponseWriter, r *http.Request, itemID uint64) *Response {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

// Get item
// (GET /i/{itemID})
func (f Facade) GetItem(w http.ResponseWriter, r *http.Request, itemID uint64) *Response {
	return GetItemJSON200Response(Item{
		DisplayID:      new(uint64),
		ExternalLink:   new(string),
		ID:             new(uint64),
		PhotoURL:       new(string),
		SocialPostLink: new(string),
		Username:       new(string),
	})
}

// Update item
// (PATCH /i/{itemID})
func (f Facade) UpdateItem(w http.ResponseWriter, r *http.Request, itemID uint64) *Response {
	return UpdateItemJSON200Response(Item{
		DisplayID:      new(uint64),
		ExternalLink:   new(string),
		ID:             new(uint64),
		PhotoURL:       new(string),
		SocialPostLink: new(string),
		Username:       new(string),
	})
}

// Get user
// (GET /{username})
func (f Facade) GetUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return GetUserJSON200Response(User{
		DisplayIDS:  []uint64{},
		Fname:       new(string),
		Lname:       new(string),
		PhotoURL:    new(string),
		SocialLinks: []string{},
		Username:    new(string),
	})
}

// Update user
// (PATCH /{username})
func (f Facade) UpdateUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return UpdateUserJSON200Response(User{
		DisplayIDS:  []uint64{},
		Fname:       new(string),
		Lname:       new(string),
		PhotoURL:    new(string),
		SocialLinks: []string{},
		Username:    new(string),
	})
}

// Create user
// (POST /{username})
func (f Facade) CreateUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return CreateUserJSON200Response(User{
		DisplayIDS:  []uint64{},
		Fname:       new(string),
		Lname:       new(string),
		PhotoURL:    new(string),
		SocialLinks: []string{},
		Username:    new(string),
	})
}
