package site

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lvlBA/online_shop/internal/management/models"

	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

func (s ServiceImpl) ListSites(ctx context.Context, req *api.ListSitesRequest) (*api.ListSitesResponse, error) {
	var pagination *models.Pagination
	if req.Pagination != nil {
		pagination = &models.Pagination{
			Page:  uint(req.Pagination.Page),
			Limit: uint(req.Pagination.Limit),
		}
	}

	sites, err := s.ctrlSite.List(ctx, &controllersSite.ListParams{
		Pagination: pagination,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "error list sites")
		s.log.Error(ctx, "failed to List site", err, "request", req)
	}

	result := make([]*api.Site, 0, len(sites))
	for i := range sites {
		result = append(result, adaptSiteToApi(sites[i]))
	}
	return &api.ListSitesResponse{Sites: result}, nil
}
