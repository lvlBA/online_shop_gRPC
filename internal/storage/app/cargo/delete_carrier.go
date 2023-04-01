package cargo

import (
	"context"
	"errors"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) DeleteCarrier(ctx context.Context, req *api.DeleteCarrierRequest) (*api.DeleteCarrierResponse, error) {
	if err := validateDeleteCargoReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlCargo.DeleteCarrier(ctx, req.Id); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found carrier")
		}
		s.log.Error(ctx, "failed to delete carrier", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete carrier")
	}

	return &api.DeleteCarrierResponse{}, nil
}

func validateDeleteCargoReq(req *api.DeleteCarrierRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
