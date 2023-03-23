package warehouse

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	controllersWarehouse "github.com/lvlBA/online_shop/internal/management/controllers/warehouse"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s *ServiceImpl) ListWarehouse(ctx context.Context, req *api.ListWarehousesRequest) (*api.ListWarehousesResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	warehouses, err := s.ctrlWarehouse.ListWarehouses(ctx, &controllersWarehouse.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		s.log.Error(ctx, "failed to List site", err, "request", req)
		return nil, status.Error(codes.Internal, "error list sites")
	}

	result := make([]*api.Warehouse, 0, len(warehouses))
	for i := range warehouses {
		result = append(result, adaptWarehouseToApi(warehouses[i]))
	}
	return &api.ListWarehousesResponse{Warehouse: result}, nil
}
