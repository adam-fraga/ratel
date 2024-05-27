package models

type Layout struct {
	Name        string
	Description string
	pages       []Page
	Path        string `default:"views/layouts"`
}

type Page struct {
	Name        string
	Description string
	Path        string `default:"views/pages"`
}

type Component struct {
	Name        string
	Description string
	Path        string `default:"views/components"`
}
