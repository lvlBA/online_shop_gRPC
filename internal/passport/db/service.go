package db

import (
	"github.com/jmoiron/sqlx"
)

type ServiceImpl struct {
	*sqlx.DB
}

func New(db *sqlx.DB) Service {
	return &ServiceImpl{
		DB: db,
	}
}

func (s *ServiceImpl) User() User {
	return &UserImpl{
		svc: s,
	}
}

func (s *ServiceImpl) Resource() Resource {
	return &ResourceImpl{
		svc: s,
	}
}

func (s *ServiceImpl) Auth() Auth {
	return &AuthImpl{
		svc: s,
	}
}
