package resource

import (
	"context"
	"errors"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (s *ServiceImpl) DeleteUserAccess(ctx context.Context,
	req *api.DeleteUserAccessRequest) (*api.DeleteUserAccessResponse, error) {
	if err := validateDeleteUserAccessReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlService.SetUserAccess(ctx, req.UserId); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to Delete User Access", err, "request", req)

		return nil, status.Error(codes.Internal, "error Delete User Access")
	}

	return &api.DeleteUserAccessResponse{}, nil
}

func validateDeleteUserAccessReq(req *api.DeleteUserAccessRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.UserId, validation.Required, is.UUIDv4),
	}.Filter()
}
