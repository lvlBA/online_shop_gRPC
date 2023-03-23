package location

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	controllersLocation "github.com/lvlBA/online_shop/internal/management/controllers/location"
	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s *ServiceImpl) ListLocation(ctx context.Context, req *api.ListLocationsRequest) (*api.ListLocationsResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	locations, err := s.ctrlLocation.ListLocation(ctx, &controllersLocation.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		s.log.Error(ctx, "failed to List location", err, "request", req)
		return nil, status.Error(codes.Internal, "error list locations")
	}

	result := make([]*api.Location, 0, len(locations))
	for i := range locations {
		result = append(result, adaptLocationToApi(locations[i]))
	}
	return &api.ListLocationsResponse{Location: result}, nil

}
