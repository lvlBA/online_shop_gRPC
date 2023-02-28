package site

import (
	"context"

	controllersSite "github.com/lvlBA/online_shop/internal/management/controllers/site"
	api "github.com/lvlBA/online_shop/pkg/management/v1"
)

type ServiceImpl struct {
	ctrlSite controllersSite.Service
	api.UnimplementedSiteServiceServer
}

func New(ctrlSite controllersSite.Service) *ServiceImpl {
	return &ServiceImpl{
		ctrlSite:                       ctrlSite,
		UnimplementedSiteServiceServer: api.UnimplementedSiteServiceServer{},
	}
}

func (s ServiceImpl) DeleteSite(ctx context.Context, request *api.DeleteSiteRequest) (*api.DeleteSiteResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) ListSites(ctx context.Context, request *api.ListSitesRequest) (*api.ListSitesResponse, error) {
	//TODO implement me
	panic("implement me")
}
