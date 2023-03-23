package location

import (
	controllersLocation "github.com/lvlBA/online_shop/internal/management/controllers/location"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type ServiceImpl struct {
	ctrlLocation controllersLocation.Service
	api.UnimplementedLocationServiceServer
	log logger.Logger
}

func New(ctrlLocation controllersLocation.Service, l logger.Logger) *ServiceImpl {
	return &ServiceImpl{
		ctrlLocation:                       ctrlLocation,
		UnimplementedLocationServiceServer: api.UnimplementedLocationServiceServer{},
		log:                                l,
	}
}
