package user

import (
	"context"

	"github.com/lvlBA/online_shop/pkg/logger"

	controllersUser "github.com/lvlBA/online_shop/internal/passport/controllers/user"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type ServiceImpl struct {
	ctrlUser controllersUser.Service
	api.UnimplementedUserServiceServer
	log logger.Logger
}

func (s ServiceImpl) GetUser(ctx context.Context, request *api.GetUserRequest) (*api.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) ChangePass(ctx context.Context, request *api.ChangePassRequest) (*api.ChangePassResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) DeleteUser(ctx context.Context, request *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) ListUsers(ctx context.Context, request *api.ListUsersRequest) (*api.ListUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func New(ctrlUser controllersUser.Service, l logger.Logger) api.UserServiceServer {
	return &ServiceImpl{
		ctrlUser:                       ctrlUser,
		UnimplementedUserServiceServer: api.UnimplementedUserServiceServer{},
		log:                            nil,
	}
}
