package database

import "github.com/google/uuid"

type Display struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Description string
	PhotoURL    string
}

type Item struct {
	ID             uuid.UUID
	UserID         uuid.UUID
	DisplayID      uuid.UUID
	ExternalLink   string
	SocialPostLink string
	PhotoURL       string
}
