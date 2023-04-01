package goods

import (
	controllerGoods "github.com/lvlBA/online_shop/internal/storage/controllers/goods"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

type ServiceImpl struct {
	ctrlGoods controllerGoods.Service
	api.UnimplementedGoodsServiceServer
	log logger.Logger
}

func New(ctrlGoods controllerGoods.Service, l logger.Logger) api.GoodsServiceServer {
	return &ServiceImpl{
		ctrlGoods:                       ctrlGoods,
		UnimplementedGoodsServiceServer: api.UnimplementedGoodsServiceServer{},
		log:                             l,
	}
}
