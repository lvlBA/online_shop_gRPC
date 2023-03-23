package models

type Location struct {
	Meta
	Name     string `db:"name"`
	SiteId   string `db:"site_id"`
	RegionId string `db:"region_id"`
}
