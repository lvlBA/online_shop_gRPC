package region

import (
	controllersRegion "github.com/lvlBA/online_shop/internal/management/controllers/region"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type ServiceImpl struct {
	ctrlRegion controllersRegion.Service
	api.UnimplementedRegionServiceServer
	log logger.Logger
}

func New(ctrlRegion controllersRegion.Service, l logger.Logger) *ServiceImpl {
	return &ServiceImpl{
		ctrlRegion:                       ctrlRegion,
		UnimplementedRegionServiceServer: api.UnimplementedRegionServiceServer{},
		log:                              l,
	}
}
