package models

type Site struct {
	Meta
	Name string `db:"name"`
}
