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

const tableNameRegion = "region"

type RegionImpl struct {
	svc service
}

type CreateRegionParams struct {
	Name   string
	SiteId string
}

func (r *RegionImpl) CreateRegion(ctx context.Context, params *CreateRegionParams) (*models.Region, error) {
	model := &models.Region{
		Meta:   models.Meta{},
		Name:   params.Name,
		SiteId: params.SiteId,
	}

	model.UpdateMeta()

	id, err := r.svc.create(ctx, tableNameRegion, model)
	fmt.Println(id)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (r *RegionImpl) GetRegion(ctx context.Context, id string) (*models.Region, error) {
	result := &models.Region{}

	query, _, err := goqu.From(tableNameRegion).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = r.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RegionImpl) DeleteRegion(ctx context.Context, id string) error {
	return r.svc.delete(ctx, tableNameRegion, id)
}

type ListRegionFilter struct {
	Pagination *models.Pagination
}

func (f *ListRegionFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

func (r *RegionImpl) ListRegion(ctx context.Context, filter *ListRegionFilter) ([]*models.Region, error) {
	ds := goqu.From(tableNameRegion).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Region, 0)
	if err = r.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
