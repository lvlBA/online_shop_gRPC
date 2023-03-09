package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	controllersUser "github.com/lvlBA/online_shop/internal/passport/controllers/user"
	"github.com/lvlBA/online_shop/internal/passport/models"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s ServiceImpl) ListUsers(ctx context.Context, req *api.ListUsersRequest) (*api.ListUsersResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	users, err := s.ctrlUser.ListUsers(ctx, &controllersUser.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "error list users")
		s.log.Error(ctx, "failed to List users", err, "request", req)
	}

	result := make([]*api.User, 0, len(users))
	for i := range users {
		result = append(result, adaptUserToApi(users[i]))
	}
	return &api.ListUsersResponse{Users: result}, nil
}
