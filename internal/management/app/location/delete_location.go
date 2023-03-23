package location

import (
	"context"
	"errors"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s *ServiceImpl) DeleteLocation(ctx context.Context, req *api.DeleteLocationRequest) (*api.DeleteLocationResponse, error) {
	if err := validateDeleteLocationReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlLocation.DeleteLocation(ctx, req.Id); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete location", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete location")
	}

	return &api.DeleteLocationResponse{}, nil
}

func validateDeleteLocationReq(req *api.DeleteLocationRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
