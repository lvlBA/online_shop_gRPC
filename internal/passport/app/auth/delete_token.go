package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s *ServiceImpl) DeleteUserToken(ctx context.Context, req *api.DeleteUserTokenRequest) (*api.DeleteUserTokenResponse, error) {
	if err := s.ctrlAuth.DeleteUserToken(ctx, req.Token); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete token", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete user")
	}

	return &api.DeleteUserTokenResponse{}, nil
}
