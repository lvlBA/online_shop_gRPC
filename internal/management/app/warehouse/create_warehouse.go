package warehouse

import (
	"context"
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	controllersWarehouse "github.com/lvlBA/online_shop/internal/management/controllers/warehouse"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s *ServiceImpl) CreateWarehouse(ctx context.Context, req *api.CreateWarehouseRequest) (*api.CreateWarehouseResponse, error) {
	if err := validateCreateWarehouseReq(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	warehouse, err := s.ctrlWarehouse.CreateWarehouse(ctx, &controllersWarehouse.CreateParams{
		Name:       req.Name,
		SiteId:     req.SiteId,
		RegionId:   req.RegionId,
		LocationId: req.LocationId,
	})
	if err != nil {
		if errors.Is(err, controllers.ErrorAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "warehouse already exists")
		}
		s.log.Error(ctx, "failed to create warehouse", err, "request", req)

		return nil, status.Error(codes.Internal, "error create warehouse")
	}

	return &api.CreateWarehouseResponse{
		Warehouse: adaptWarehouseToApi(warehouse),
	}, nil
}

func validateCreateWarehouseReq(req *api.CreateWarehouseRequest) error {
	return validation.Errors{
		"name": validation.Validate(req.Name, validation.Required),
	}.Filter()
}

func adaptWarehouseToApi(model *models.Warehouse) *api.Warehouse {
	return &api.Warehouse{
		Id:   model.ID,
		Name: model.Name,
	}
}
