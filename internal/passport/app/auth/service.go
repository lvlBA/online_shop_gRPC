package auth

import (
	"github.com/lvlBA/online_shop/pkg/logger"

	controllerAuth "github.com/lvlBA/online_shop/internal/passport/controllers/auth"
	controllerResource "github.com/lvlBA/online_shop/internal/passport/controllers/resource"
	controllerUser "github.com/lvlBA/online_shop/internal/passport/controllers/user"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type ServiceImpl struct {
	api.UnimplementedAuthServiceServer
	ctrlAuth     controllerAuth.Service
	ctrlResource controllerResource.Service
	ctrlUser     controllerUser.Service
	log          logger.Logger
}

type Config struct {
	log          logger.Logger
	ctrlAuth     controllerAuth.Service
	ctrlResource controllerResource.Service
	ctrlUser     controllerUser.Service
}

func New(cfg *Config) api.AuthServiceServer {
	return &ServiceImpl{
		ctrlAuth:     cfg.ctrlAuth,
		ctrlResource: cfg.ctrlResource,
		ctrlUser:     cfg.ctrlUser,
		log:          cfg.log,
	}
}
