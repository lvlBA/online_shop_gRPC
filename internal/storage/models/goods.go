package models

type Goods struct {
	Meta
	Name   string  `db:"name"`
	Weight int     `db:"weight"`
	Length int     `db:"length"`
	Width  int     `db:"width"`
	Height int     `db:"height"`
	Price  float64 `db:"price"`
}
