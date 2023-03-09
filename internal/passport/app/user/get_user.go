package user

import (
	"context"
	"errors"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
)

func (s ServiceImpl) GetUser(ctx context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	if err := validateGetUserReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.ctrlUser.GetUser(ctx, req.Id)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "site not found")
		}
		s.log.Error(ctx, "failed to get site", err, "request", req)

		return nil, status.Error(codes.Internal, "error get site")
	}

	return &api.GetUserResponse{
		User: adaptUserToApi(user),
	}, nil
}

func validateGetUserReq(req *api.GetUserRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
