package site

import (
	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type ServiceImpl struct {
	ctrlSite controllersSite.Service
	api.UnimplementedSiteServiceServer
	log logger.Logger
}

func New(ctrlSite controllersSite.Service, l logger.Logger) *ServiceImpl {
	return &ServiceImpl{
		ctrlSite:                       ctrlSite,
		UnimplementedSiteServiceServer: api.UnimplementedSiteServiceServer{},
		log:                            l,
	}
}
