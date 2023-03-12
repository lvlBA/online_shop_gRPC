package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllerAuth "github.com/lvlBA/online_shop/internal/passport/controllers/auth"
	controllerUser "github.com/lvlBA/online_shop/internal/passport/controllers/user"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s *ServiceImpl) GetUserToken(ctx context.Context, req *api.GetUserTokenRequest) (*api.GetUserTokenResponse, error) {
	// FIXME: валидация

	user, err := s.ctrlUser.GetUser(ctx, &controllerUser.GetUserParams{
		Login:    &req.Login,
		Password: &req.Password,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.Unauthenticated, "invalid user or password")
		}

		return nil, status.Error(codes.Internal, "error get user token")
	}

	token, err := s.ctrlAuth.CreateUserToken(ctx, &controllerAuth.CreateUserTokenRequest{
		UserID: user.ID,
	})
	if err != nil {
		s.log.Error(ctx, "failed to get token", err, "request", req)

		return nil, status.Error(codes.Internal, "error get token")
	}

	return &api.GetUserTokenResponse{
		Token: token.Token,
	}, nil
}
