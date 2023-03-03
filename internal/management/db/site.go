package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/lvlBA/online_shop/internal/management/models"
	utilspagination "github.com/lvlBA/online_shop/pkg/utils/pagination"
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

	query, _, err := goqu.From(tableNameSite).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = s.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *SiteImpl) DeleteSite(ctx context.Context, id string) error {
	return s.svc.delete(ctx, tableNameSite, id)
}

type ListSitesFilter struct {
	Pagination *models.Pagination
}

func (f *ListSitesFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

func (s *SiteImpl) ListSites(ctx context.Context, filter *ListSitesFilter) ([]*models.Site, error) {
	ds := goqu.From(tableNameSite).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Site, 0)
	if err = s.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
