package models

type OrdersStore struct {
	Meta
	Name        string `db:"name"`
	SiteId      string `db:"site_id"`
	RegionId    string `db:"region_id"`
	LocationId  string `db:"location_id"`
	WarehouseId string `db:"warehouse_id"`
}
