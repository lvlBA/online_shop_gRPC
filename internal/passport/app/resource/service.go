package resource

import (
	controllersservice "github.com/lvlBA/online_shop/internal/passport/controllers/resource"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type ServiceImpl struct {
	ctrlService controllersservice.Service
	api.UnimplementedResourceServiceServer
	log logger.Logger
}

func New(ctrlService controllersservice.Service, l logger.Logger) api.ResourceServiceServer {
	return &ServiceImpl{
		ctrlService:                        ctrlService,
		UnimplementedResourceServiceServer: api.UnimplementedResourceServiceServer{},
		log:                                l,
	}
}
