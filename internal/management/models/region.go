package models

type Region struct {
	Meta
	Name   string `db:"name"`
	SiteId string `db:"site_id"`
}
