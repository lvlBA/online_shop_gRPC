package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/lvlBA/online_shop/internal/passport/models"
	utilspagination "github.com/lvlBA/online_shop/pkg/utils/pagination"
)

type CreateServiceParams struct {
	Urn string
}

type ResourceImplementation struct {
	svc service
}

type ListServiceFilter struct {
	Pagination *models.Pagination
}

func (f *ListServiceFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

const tableNameService = "service"

func (r *ResourceImplementation) CreateService(ctx context.Context, params *CreateServiceParams) (*models.Resource,
	error) {
	model := &models.Resource{
		Meta: models.Meta{},
		Urn:  params.Urn,
	}
	model.UpdateMeta()

	id, err := r.svc.create(ctx, tableNameService, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (r *ResourceImplementation) GetResource(ctx context.Context, id string) (*models.Resource, error) {
	result := &models.Resource{}

	query, _, err := goqu.From(tableNameService).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
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

func (r *ResourceImplementation) DeleteResource(ctx context.Context, id string) error {
	return r.svc.delete(ctx, tableNameService, id)
}

func (r *ResourceImplementation) ListResource(ctx context.Context, filter *ListServiceFilter) ([]*models.Resource,
	error) {
	ds := goqu.From(tableNameService).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Resource, 0)
	if err = r.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *ResourceImplementation) SetUserAccess(ctx context.Context, id string) error {

}

func (r *ResourceImplementation) DeleteUserAccess(ctx context.Context, id string) error {

}
