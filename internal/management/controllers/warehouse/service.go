package warehouse

import "github.com/lvlBA/online_shop/internal/management/db"

type ServiceImpl struct {
	db db.Service
}

func New(db db.Service) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
