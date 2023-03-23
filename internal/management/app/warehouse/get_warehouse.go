package warehouse

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

func (s *ServiceImpl) GetWarehouse(ctx context.Context, req *api.GetWarehouseRequest) (*api.GetWarehouseResponse, error) {
	if err := validateGetWarehouseReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	warehouse, err := s.ctrlWarehouse.GetWarehouse(ctx, req.Id)
	if err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "warehouse not found")
		}
		s.log.Error(ctx, "failed to get warehouse", err, "request", req)

		return nil, status.Error(codes.Internal, "error get warehouse")
	}

	return &api.GetWarehouseResponse{
		Warehouse: adaptWarehouseToApi(warehouse),
	}, nil
}

func validateGetWarehouseReq(req *api.GetWarehouseRequest) error {
	return validation.Errors{
		"id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
