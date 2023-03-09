package user

import (
	"github.com/lvlBA/online_shop/pkg/logger"

	controllersUser "github.com/lvlBA/online_shop/internal/passport/controllers/user"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type ServiceImpl struct {
	ctrlUser controllersUser.Service
	api.UnimplementedUserServiceServer
	log logger.Logger
}

func New(ctrlUser controllersUser.Service, l logger.Logger) api.UserServiceServer {
	return &ServiceImpl{
		ctrlUser:                       ctrlUser,
		UnimplementedUserServiceServer: api.UnimplementedUserServiceServer{},
		log:                            l,
	}
}
