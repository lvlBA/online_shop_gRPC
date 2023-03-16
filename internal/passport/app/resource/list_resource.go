package resource

import (
	"context"
	controllersResource "github.com/lvlBA/online_shop/internal/passport/controllers/resource"
	"github.com/lvlBA/online_shop/internal/passport/models"
	api "github.com/lvlBA/online_shop/pkg/passport/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServiceImpl) ListResource(ctx context.Context, req *api.ListResourceRequest) (*api.ListResourceResponse,
	error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	resources, err := s.ctrlService.ListResource(ctx, &controllersResource.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		s.log.Error(ctx, "failed to List resources", err, "request", req)
		return nil, status.Error(codes.Internal, "error list resources")
	}

	result := make([]*api.Resource, 0, len(resources))
	for i := range resources {
		result = append(result, adaptResourceToApi(resources[i]))
	}
	return &api.ListResourceResponse{Resource: result}, nil
}
