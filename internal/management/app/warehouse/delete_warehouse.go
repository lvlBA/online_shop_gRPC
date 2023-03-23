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

func (s *ServiceImpl) DeleteWarehouse(ctx context.Context, req *api.DeleteWarehouseRequest) (*api.DeleteWarehouseResponse, error) {
	if err := validateDeleteWarehouseReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.ctrlWarehouse.DeleteWarehouse(ctx, req.Id); err != nil {
		if errors.Is(err, controllers.ErrorNotFound) {
			return nil, status.Error(codes.NotFound, "Not found")
		}
		s.log.Error(ctx, "failed to delete warehouse", err, "request", req)

		return nil, status.Error(codes.Internal, "error delete warehouse")
	}

	return &api.DeleteWarehouseResponse{}, nil
}

func validateDeleteWarehouseReq(req *api.DeleteWarehouseRequest) error {
	return validation.Errors{
		"Id": validation.Validate(req.Id, validation.Required, is.UUIDv4),
	}.Filter()
}
