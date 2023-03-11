package models

type Resource struct {
	Meta
	Urn    string `db:"urn"`
	Access bool   `db:"access"`
}
