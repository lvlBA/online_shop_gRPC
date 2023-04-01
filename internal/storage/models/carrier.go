package models

type Carrier struct {
	Meta
	Name         string  `db:"name"`
	Capacity     int     `db:"capacity"`
	Price        float64 `db:"price"`
	Availability bool    `db:"availability"`
}
