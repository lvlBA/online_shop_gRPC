package cargo

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	controllerCargo "github.com/lvlBA/online_shop/internal/storage/controllers/cargo"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) GetCarrier(ctx context.Context, req *api.GetCarrierRequest) (*api.GetCarrierResponse, error) {
	if err := validateGetCarrierReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	carrier, err := s.ctrlCargo.GetCarrier(ctx, &controllerCargo.GetCargoParams{
		ID:   &req.Id,
		Name: &req.Name,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "carrier not found")
		}
		s.log.Error(ctx, "failed to get carrier", err, "request", req)

		return nil, status.Error(codes.Internal, "error get carrier")
	}

	return &api.GetCarrierResponse{
		Carrier: adaptCarrierToApi(carrier),
	}, nil

}

func validateGetCarrierReq(req *api.GetCarrierRequest) error {
	return validation.Errors{
		"id":   validation.Validate(req.Id, validation.Required, is.UUIDv4),
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}
