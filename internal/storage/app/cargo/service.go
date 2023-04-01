package cargo

import (
	controllerCargo "github.com/lvlBA/online_shop/internal/storage/controllers/cargo"
	"github.com/lvlBA/online_shop/pkg/logger"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

type ServiceImpl struct {
	ctrlCargo controllerCargo.Service
	api.UnimplementedCargoServiceServer
	log logger.Logger
}

func New(ctrlCargo controllerCargo.Service, l logger.Logger) api.CargoServiceServer {
	return &ServiceImpl{
		ctrlCargo:                       ctrlCargo,
		UnimplementedCargoServiceServer: api.UnimplementedCargoServiceServer{},
		log:                             l,
	}
}
