package cargo

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	controllersCarrier "github.com/lvlBA/online_shop/internal/storage/controllers/cargo"
	"github.com/lvlBA/online_shop/internal/storage/models"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) CreateCarrier(ctx context.Context, req *api.CreateCarrierRequest) (*api.CreateCarrierResponse, error) {
	if err := validateCreateCarrierReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	carrier, err := s.ctrlCargo.CreateCarrier(ctx, &controllersCarrier.CreateParams{
		Name:         req.Name,
		Capacity:     int(req.Capacity),
		Price:        float64(req.Price),
		Availability: false,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "goods already exists")
		}
		s.log.Error(ctx, "failed to create goods", err, "request", req)

		return nil, status.Error(codes.Internal, "error create goods")
	}

	return &api.CreateCarrierResponse{
		Carrier: adaptCarrierToApi(carrier),
	}, nil
}

func validateCreateCarrierReq(req *api.CreateCarrierRequest) error {
	return validation.Errors{
		"name":         validation.Validate(req.Name, validation.Required),
		"capacity":     validation.Validate(req.Capacity, validation.Required),
		"price":        validation.Validate(req.Price, validation.Required),
		"availability": validation.Validate(req.Availability, validation.Required),
	}.Filter()
}

func adaptCarrierToApi(model *models.Carrier) *api.Carrier {
	return &api.Carrier{
		Id:           model.ID,
		Name:         model.Name,
		Capacity:     uint64(model.Capacity),
		Price:        float32(model.Price),
		Availability: model.Availability,
	}
}
