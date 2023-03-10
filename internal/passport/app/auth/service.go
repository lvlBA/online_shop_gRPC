package auth

import (
	"context"

	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

type ServiceImpl struct {
}

func (s *ServiceImpl) GetUserToken(ctx context.Context, in *api.GetUserTokenRequest) (*api.GetUserTokenResponse, error) {
}
func (s *ServiceImpl) CheckUser(ctx context.Context, in *api.CheckUserRequest) (*api.CheckUserResponse, error) {
}
