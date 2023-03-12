package auth

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s *ServiceImpl) CheckUser(ctx context.Context, req *api.CheckUserRequest) (*api.CheckUserResponse, error) {
	if _, err := s.ctrlUser.CheckUser(ctx, req.Resource); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "resource not found")
		}
		s.log.Error(ctx, "failed to get user", err, "request", req)

		return nil, status.Error(codes.Internal, "error check user")
	}

	return &api.CheckUserResponse{}, nil
}
