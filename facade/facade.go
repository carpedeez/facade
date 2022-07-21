package facade

import (
	"database/sql"
	"net/http"
)

type Facade struct {
	DB *sql.DB
}

// Upload file
// (POST /assets)
func (f Facade) UploadFile(w http.ResponseWriter, r *http.Request) {
}

// Get user
// (GET /{username})
func (f Facade) GetUser(w http.ResponseWriter, r *http.Request, username string)

// Update user
// (PATCH /{username})
func (f Facade) UpdateUser(w http.ResponseWriter, r *http.Request, username string)

// Create user
// (POST /{username})
func (f Facade) CreateUser(w http.ResponseWriter, r *http.Request, username string)

// Delete display
// (DELETE /{username}/{displayID})
func (f Facade) DeleteDisplay(w http.ResponseWriter, r *http.Request, username string, displayID uint64)

// Get display
// (GET /{username}/{displayID})
func (f Facade) GetDisplays(w http.ResponseWriter, r *http.Request, username string, displayID uint64)

// Update display
// (PATCH /{username}/{displayID})
func (f Facade) UpdateDisplay(w http.ResponseWriter, r *http.Request, username string, displayID uint64)

// Create display
// (POST /{username}/{displayID})
func (f Facade) CreateDisplay(w http.ResponseWriter, r *http.Request, username string, displayID uint64)

// Delete item
// (DELETE /{username}/{displayID}/{itemID})
func (f Facade) DeleteItem(w http.ResponseWriter, r *http.Request, username string, displayID uint64, itemID uint64)

// Get item
// (GET /{username}/{displayID}/{itemID})
func (f Facade) GetItem(w http.ResponseWriter, r *http.Request, username string, displayID uint64, itemID uint64)

// Update item
// (PATCH /{username}/{displayID}/{itemID})
func (f Facade) UpdateItem(w http.ResponseWriter, r *http.Request, username string, displayID uint64, itemID uint64)

// Create item
// (POST /{username}/{displayID}/{itemID})
func (f Facade) CreateItem(w http.ResponseWriter, r *http.Request, username string, displayID uint64, itemID uint64)
