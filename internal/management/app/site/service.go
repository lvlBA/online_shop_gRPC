package site

import (
	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type ServiceImpl struct {
	ctrlSite controllersSite.Service
	api.UnimplementedSiteServiceServer
}

func New(ctrlSite controllersSite.Service) *ServiceImpl {
	return &ServiceImpl{
		ctrlSite:                       ctrlSite,
		UnimplementedSiteServiceServer: api.UnimplementedSiteServiceServer{},
	}
}
