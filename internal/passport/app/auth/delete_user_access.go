package auth

import (
	"context"
	"errors"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServiceImpl) DeleteUserAccess(
	ctx context.Context, req *api.DeleteUserAccessRequest) (*api.DeleteUserAccessResponse, error) {

	if err := validateDeleteAccessReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlAuth.DeleteUserAccess(ctx, &req.UserId, &req.ResourceId); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete access", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete access")
	}

	return &api.DeleteUserAccessResponse{}, nil
}

func validateDeleteAccessReq(req *api.DeleteUserAccessRequest) error {
	return validation.Errors{
		"UserId": validation.Validate(
			req.UserId,
			validation.Required,
			is.UUIDv4),
		"ResourceId": validation.Validate(
			req.ResourceId,
			validation.Required,
			is.UUIDv4),
	}.Filter()
}
