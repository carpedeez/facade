package facade

import (
	"log"
	"net/http"

	"github.com/carpedeez/store/database"
)

type Facade struct {
	Querier *database.Querier
}

// Upload file
// (POST /assets)
func (f Facade) UploadFile(w http.ResponseWriter, r *http.Request) {}

// Create display
// (POST /d)
func (f Facade) CreateDisplay(w http.ResponseWriter, r *http.Request, params CreateDisplayParams) {
	log.Printf("%+v", *r)
	// id, err := f.Querier.CreateDisplay(params.Display.Username, params.Display.Title, params.Display.Description)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(id)
}

// Delete display
// (DELETE /d/{displayID})
func (f Facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) {

}

// Get display
// (GET /d/{displayID})
func (f Facade) GetDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) {}

// Update display
// (PATCH /d/{displayID})
func (f Facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, displayID uint64) {}

// Create item
// (POST /i)
func (f Facade) CreateItem(w http.ResponseWriter, r *http.Request, params CreateItemParams) {}

// Delete item
// (DELETE /i/{itemID})
func (f Facade) DeleteItem(w http.ResponseWriter, r *http.Request, itemID uint64) {}

// Get item
// (GET /i/{itemID})
func (f Facade) GetItem(w http.ResponseWriter, r *http.Request, itemID uint64) {}

// Update item
// (PATCH /i/{itemID})
func (f Facade) UpdateItem(w http.ResponseWriter, r *http.Request, itemID uint64) {}

// Get user
// (GET /{username})
func (f Facade) GetUser(w http.ResponseWriter, r *http.Request, username string) {}

// Update user
// (PATCH /{username})
func (f Facade) UpdateUser(w http.ResponseWriter, r *http.Request, username string) {}

// Create user
// (POST /{username})
func (f Facade) CreateUser(w http.ResponseWriter, r *http.Request, username string) {}
