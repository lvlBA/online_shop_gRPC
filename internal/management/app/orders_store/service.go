package orders_store

import (
	"github.com/lvlBA/online_shop/pkg/logger"

	controllersOrdersStore "github.com/lvlBA/online_shop/internal/management/controllers/orders_store"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type ServiceImpl struct {
	ctrlOrdersStore controllersOrdersStore.Service
	api.UnimplementedOrdersStoreServiceServer
	log logger.Logger
}

func New(ctrlOrdersStore controllersOrdersStore.Service, l logger.Logger) *ServiceImpl {
	return &ServiceImpl{
		ctrlOrdersStore:                       ctrlOrdersStore,
		UnimplementedOrdersStoreServiceServer: api.UnimplementedOrdersStoreServiceServer{},
		log:                                   l,
	}
}
