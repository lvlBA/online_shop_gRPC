package db

import (
	"context"
	"fmt"
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
	result := &models.Site{}

	query, _, err := goqu.From(tableNameSite).Select("*").Where(goqu.Ex{"id": id}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = s.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}

	return result, nil
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
