package db

import (
	sql "github.com/jmoiron/sqlx"
)

type ServiceImpl struct {
	sqlClient *sql.DB
}

func New(sqlClient *sql.DB) *ServiceImpl {
	return &ServiceImpl{
		sqlClient: sqlClient,
	}
}

func (s *ServiceImpl) Site() Site {
	return &SiteImpl{
		svc: s,
	}
}
