package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	utilsPagination "github.com/lvlBA/online_shop/pkg/utils/pagination"

	"github.com/doug-martin/goqu/v9"
	"github.com/lvlBA/online_shop/internal/storage/models"
)

const tableNameCarrier = "carrier"

type cargoImpl struct {
	svc sqlService
}

type CreateCargoParams struct {
	Name         string
	Capacity     int
	Price        float64
	Availability bool
}

func (c *cargoImpl) CreateCarrier(ctx context.Context, params *CreateCargoParams) (*models.Carrier, error) {
	model := &models.Carrier{
		Meta:         models.Meta{},
		Name:         params.Name,
		Capacity:     params.Capacity,
		Price:        params.Price,
		Availability: params.Availability,
	}
	model.UpdateMeta()

	id, err := c.svc.create(ctx, tableNameCarrier, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

type GetCargoParams struct {
	UserID *string
	Name   *string
}

func (p *GetCargoParams) filter(sd *goqu.SelectDataset) (*goqu.SelectDataset, error) {
	switch {
	case p.UserID != nil:
		return sd.Where(goqu.Ex{"id": *p.UserID}), nil
	case p.Name != nil:
		return sd.Where(goqu.Ex{"name": p.Name}), nil
	default:
		return nil, errors.New("undefined behavior: id is not set and name is not set")
	}
}

func (c *cargoImpl) GetCarrier(ctx context.Context, params *GetCargoParams) (*models.Carrier, error) {
	sd, err := params.filter(goqu.From(tableNameCarrier).Select("*"))
	if err != nil {
		return nil, fmt.Errorf("failed to create filter: %w", err)
	}

	query, _, err := sd.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	result := &models.Carrier{}
	if err = c.svc.GetContext(ctx, result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}

		return nil, err
	}

	return result, nil
}

func (c *cargoImpl) DeleteCarrier(ctx context.Context, id string) error {
	return c.svc.delete(ctx, tableNameCarrier, id)
}

type ListCargoFilter struct {
	Pagination *models.Pagination
}

func (f *ListCargoFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilsPagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

func (c *cargoImpl) ListCarrier(ctx context.Context, filter *ListCargoFilter) ([]*models.Carrier, error) {
	ds := goqu.From(tableNameCarrier).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Carrier, 0)
	if err = c.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
