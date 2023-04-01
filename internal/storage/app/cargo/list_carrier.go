package cargo

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	controllerCargo "github.com/lvlBA/online_shop/internal/storage/controllers/cargo"
	"github.com/lvlBA/online_shop/internal/storage/models"
	api "github.com/lvlBA/online_shop/pkg/storage/v1"
)

func (s *ServiceImpl) ListCarriers(ctx context.Context, req *api.ListCarrierRequest) (*api.ListCarrierResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	carriers, err := s.ctrlCargo.ListCarrier(ctx, &controllerCargo.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		s.log.Error(ctx, "failed to List carriers", err, "request", req)
		return nil, status.Error(codes.Internal, "error list carriers")
	}

	result := make([]*api.Carrier, 0, len(carriers))
	for i := range carriers {
		result = append(result, adaptCarrierToApi(carriers[i]))
	}
	return &api.ListCarrierResponse{Carrier: result}, nil
}
