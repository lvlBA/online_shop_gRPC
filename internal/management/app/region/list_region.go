package region

import (
	"context"
	controllersRegion "github.com/lvlBA/online_shop/internal/management/controllers/region"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/models"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s ServiceImpl) ListRegion(ctx context.Context, req *api.ListRegionsRequest) (*api.ListRegionsResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	regions, err := s.ctrlRegion.ListRegion(ctx, &controllersRegion.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		s.log.Error(ctx, "failed to List region", err, "request", req)
		return nil, status.Error(codes.Internal, "error list regions")
	}

	result := make([]*api.Region, 0, len(regions))
	for i := range regions {
		result = append(result, adaptRegionToApi(regions[i]))
	}
	return &api.ListRegionsResponse{Region: result}, nil
}
