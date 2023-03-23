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

func (s *ServiceImpl) GetLocation(ctx context.Context, req *api.GetLocationRequest) (*api.GetLocationResponse, error) {
	if err := validateGetLocationReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	location, err := s.ctrlLocation.GetLocation(ctx, req.Id)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "location not found")
		}
		s.log.Error(ctx, "failed to get location", err, "request", req)

		return nil, status.Error(codes.Internal, "error get location")
	}

	return &api.GetLocationResponse{
		Location: adaptLocationToApi(location),
	}, nil
}

func validateGetLocationReq(req *api.GetLocationRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
