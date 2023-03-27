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

const tableNameWarehouse = "warehouse"

type warehouseImpl struct {
	svc sqlService
}

type CreateWarehouseParams struct {
	Name       string
	SiteId     string
	RegionId   string
	LocationId string
}

func (w *warehouseImpl) CreateWarehouse(ctx context.Context, params *CreateWarehouseParams) (*models.Warehouse, error) {
	model := &models.Warehouse{
		Meta:       models.Meta{},
		Name:       params.Name,
		SiteId:     params.SiteId,
		RegionId:   params.RegionId,
		LocationId: params.LocationId,
	}
	model.UpdateMeta()

	id, err := w.svc.create(ctx, tableNameWarehouse, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (w *warehouseImpl) GetWarehouse(ctx context.Context, id string) (*models.Warehouse, error) {
	result := &models.Warehouse{}

	query, _, err := goqu.From(tableNameWarehouse).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = w.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (w *warehouseImpl) DeleteWarehouse(ctx context.Context, id string) error {
	return w.svc.delete(ctx, tableNameWarehouse, id)
}

type ListWareHouseFilter struct {
	Pagination *models.Pagination
}

func (f *ListWareHouseFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}
	return ds
}

func (w *warehouseImpl) ListWarehouses(ctx context.Context, filter *ListWareHouseFilter) ([]*models.Warehouse, error) {
	ds := goqu.From(tableNameWarehouse).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Warehouse, 0)
	if err = w.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
