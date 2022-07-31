package facade

type CreateDisplayParams struct {
	Display struct {
		Username    string
		Title       string
		Description string
	}
}

type CreateItemParams struct {
	Item struct {
		ExternalLink string
	}
}
