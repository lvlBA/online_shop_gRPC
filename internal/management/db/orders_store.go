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

const tableNameOrderStore = "orders_store"

type OrdersStoreImpl struct {
	svc service
}

type CreateOrdersStoreParams struct {
	Name        string
	SiteId      string
	RegionId    string
	LocationId  string
	WarehouseId string
}

func (o *OrdersStoreImpl) CreateOrderStore(ctx context.Context, params *CreateOrdersStoreParams) (*models.OrdersStore, error) {
	model := &models.OrdersStore{
		Meta:        models.Meta{},
		Name:        params.Name,
		SiteId:      params.SiteId,
		RegionId:    params.RegionId,
		LocationId:  params.LocationId,
		WarehouseId: params.WarehouseId,
	}
	model.UpdateMeta()

	id, err := o.svc.create(ctx, tableNameOrderStore, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (o *OrdersStoreImpl) GetOrderStore(ctx context.Context, id string) (*models.OrdersStore, error) {
	result := &models.OrdersStore{}

	query, _, err := goqu.From(tableNameOrderStore).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = o.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (o *OrdersStoreImpl) DeleteOrderStore(ctx context.Context, id string) error {
	return o.svc.delete(ctx, tableNameOrderStore, id)
}

type ListOrdersStoreFilter struct {
	Pagination *models.Pagination
}

func (f *ListOrdersStoreFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}
	return ds
}

func (o *OrdersStoreImpl) ListOrderStores(ctx context.Context, filter *ListOrdersStoreFilter) ([]*models.OrdersStore, error) {
	ds := goqu.From(tableNameOrderStore).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.OrdersStore, 0)
	if err = o.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
