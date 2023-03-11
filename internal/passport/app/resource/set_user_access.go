package resource

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

func (s *ServiceImpl) SetUserAccess(ctx context.Context,
	req *api.SetUserAccessRequest) (*api.SetUserAccessResponse, error) {
	if err := validateSetUserAccessReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlService.SetUserAccess(ctx, req.UserId); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to Set User Access", err, "request", req)

		return nil, status.Error(codes.Internal, "error Set User Access")
	}

	return &api.SetUserAccessResponse{}, nil
}

func validateSetUserAccessReq(req *api.SetUserAccessRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.UserId, validation.Required, is.UUIDv4),
	}.Filter()
}
