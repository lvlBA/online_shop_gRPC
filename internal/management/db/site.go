package db

import (
	"context"
	"github.com/doug-martin/goqu"

	"github.com/lvlBA/online_shop/internal/management/models"
)

const tableNameSite = "sites"

type SiteImpl struct {
	svc service
}

type CreateSiteParams struct {
	Name string
}

func (s *SiteImpl) CreateSite(ctx context.Context, params *CreateSiteParams) (*models.Site, error) {
	model := &models.Site{
		Meta: models.Meta{},
		Name: params.Name,
	}
	model.UpdateMeta()

	id, err := s.svc.create(ctx, tableNameSite, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (s *SiteImpl) GetSite(ctx context.Context, id string) (*models.Site, error) {
	panic("unimplemented")
}

func (s *SiteImpl) DeleteSite(ctx context.Context, id string) error {
	panic("unimplemented")
}

type ListSitesFilter struct {
}

func (f *ListSitesFilter) Filter(ds *goqu.Dataset) *goqu.Dataset {
	// TODO: implements

	return ds
}

func (s *SiteImpl) ListSites(ctx context.Context, filter *ListSitesFilter) ([]*models.Site, error) {
	panic("unimplemented")
}
