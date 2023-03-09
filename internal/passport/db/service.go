package db

import (
	"github.com/jmoiron/sqlx"
)

type ServiceImpl struct {
	*sqlx.DB
}

func New(db *sqlx.DB) *ServiceImpl {
	return &ServiceImpl{
		DB: db,
	}
}
