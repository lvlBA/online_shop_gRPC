package goods

import (
	"github.com/lvlBA/online_shop/internal/storage/db"
)

type ServiceImpl struct {
	db db.Service
}

func New(db db.Service) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
