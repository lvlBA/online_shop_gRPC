package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	utilspagination "github.com/lvlBA/online_shop/pkg/utils/pagination"

	"github.com/lvlBA/online_shop/internal/management/models"
)

const tableNameLocation = "location"

type LocationImpl struct {
	svc service
}

type CreateLocationParams struct {
	Name string
}

func (l *LocationImpl) CreateLocation(ctx context.Context, params *CreateLocationParams) (*models.Location, error) {
	model := &models.Location{
		Meta: models.Meta{},
		Name: params.Name,
	}
	model.UpdateMeta()

	id, err := l.svc.create(ctx, tableNameLocation, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

func (l *LocationImpl) GetLocation(ctx context.Context, id string) (*models.Location, error) {
	result := &models.Location{}

	query, _, err := goqu.From(tableNameLocation).Select("*").Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	if err = l.svc.GetContext(ctx, result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (l *LocationImpl) DeleteLocation(ctx context.Context, id string) error {
	return l.svc.delete(ctx, tableNameLocation, id)
}

type ListLocationFilter struct {
	Pagination *models.Pagination
}

func (f *ListLocationFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}
	return ds
}

func (l *LocationImpl) ListLocation(ctx context.Context, filter *ListLocationFilter) ([]*models.Location, error) {
	ds := goqu.From(tableNameLocation).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Location, 0)
	if err = l.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
