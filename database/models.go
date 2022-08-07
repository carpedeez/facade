package database

type Display struct {
	ID          int64
	UserID      string
	Title       string
	Description string
	PhotoURL    string
}

type Item struct {
	ID             int64
	UserID         string
	DisplayID      int64
	ExternalLink   string
	SocialPostLink string
	PhotoURL       string
}
