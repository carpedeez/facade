package facade

import (
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
func (f Facade) CreateDisplay(w http.ResponseWriter, r *http.Request, params CreateDisplayParams) *Response {
	if params.Display.Description == nil || params.Display.Title == nil || params.Display.Username == nil { // dereferencing nil pointers = panic, so we gotta check before we call CreateDisplay
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}
	id, err := f.Querier.CreateDisplay(*params.Display.Username, *params.Display.Title, *params.Display.Description) // should we respond with the serial number and redirect them? should we respond with the whole object?
	if err != nil {
		return CreateDisplayJSONDefaultResponse(Error{
			Code:    http.StatusBadRequest, // this is placeholder, we need to figure out what errors we should respond with. also, is the error object (as opposed to normal status codes) ideal?
			Message: err.Error(),
		})
	}

	return CreateDisplayJSON200Response(Display{
		Description: new(string), // this weird need for a pointer to a string is a result of not making everything required. it has to do with go's `omitempty` json flag.
		ItemIDS:     []int64{},   // ideally, we should create types for every single request and response and make everything required. lots of time tho
		ID:          &id,
		PhotoURL:    new(string),
		Title:       new(string),
		Username:    new(string),
	})
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
		Description: new(string),
		ID:          new(int64),
		ItemIDS:     []int64{},
		PhotoURL:    new(string),
		Title:       new(string),
		Username:    new(string),
	})
}

// Update display
// (PATCH /d/{displayID})
func (f Facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID int64) *Response {
	return UpdateDisplayJSON200Response(Display{
		Description: new(string),
		ID:          new(int64),
		ItemIDS:     []int64{},
		PhotoURL:    new(string),
		Title:       new(string),
		Username:    new(string),
	})
}

// Create Item
// (POST /i)
func (f Facade) CreateItem(w http.ResponseWriter, r *http.Request, params CreateItemParams) *Response {
	return CreateItemJSON200Response(Item{
		DisplayID:      new(int64),
		ExternalLink:   new(string),
		ID:             new(int64),
		PhotoURL:       new(string),
		SocialPostLink: new(string),
		Username:       new(string),
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
		DisplayID:      new(int64),
		ExternalLink:   new(string),
		ID:             new(int64),
		PhotoURL:       new(string),
		SocialPostLink: new(string),
		Username:       new(string),
	})
}

// Update item
// (PATCH /i/{itemID})
func (f Facade) UpdateItem(w http.ResponseWriter, r *http.Request, itemID int64) *Response {
	return UpdateItemJSON200Response(Item{
		DisplayID:      new(int64),
		ExternalLink:   new(string),
		ID:             new(int64),
		PhotoURL:       new(string),
		SocialPostLink: new(string),
		Username:       new(string),
	})
}

// Get user
// (GET /{username})
func (f Facade) GetUser(w http.ResponseWriter, r *http.Request, username string) *Response {
	return GetUserJSON200Response(User{
		DisplayIDS:  []int64{},
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
		DisplayIDS:  []int64{},
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
		DisplayIDS:  []int64{},
		Fname:       new(string),
		Lname:       new(string),
		PhotoURL:    new(string),
		SocialLinks: []string{},
		Username:    new(string),
	})
}
