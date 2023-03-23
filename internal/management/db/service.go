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

func (s *ServiceImpl) Site() Site {
	return &SiteImpl{
		svc: s,
	}
}

func (s *ServiceImpl) Location() Location {
	return &LocationImpl{
		svc: s,
	}
}

func (s *ServiceImpl) Region() Region {
	return &RegionImpl{
		svc: s,
	}
}

func (s *ServiceImpl) Warehouse() Warehouse {
	return &WarehouseImpl{
		svc: s,
	}
}

func (s *ServiceImpl) OrdersStore() OrdersStore {
	return &OrdersStoreImpl{
		svc: s,
	}
}
