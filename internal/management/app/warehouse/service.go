package warehouse

import (
	controllersWarehouse "github.com/lvlBA/online_shop/internal/management/controllers/warehouse"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type ServiceImpl struct {
	ctrlWarehouse controllersWarehouse.Service
	api.UnimplementedWarehouseServiceServer
	log logger.Logger
}

func New(ctrlWarehouse controllersWarehouse.Service, l logger.Logger) *ServiceImpl {
	return &ServiceImpl{
		ctrlWarehouse:                       ctrlWarehouse,
		UnimplementedWarehouseServiceServer: api.UnimplementedWarehouseServiceServer{},
		log:                                 l,
	}
}
