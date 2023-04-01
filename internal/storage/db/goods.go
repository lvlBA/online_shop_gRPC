package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"

	"github.com/lvlBA/online_shop/internal/storage/models"
	utilsPagination "github.com/lvlBA/online_shop/pkg/utils/pagination"
)

const tableNameGoods = "goods"

type goodsImpl struct {
	svc sqlService
}

type CreateGoodsParams struct {
	Name   string
	Weight int
	Length int
	Width  int
	Height int
	Price  float64
}

func (g *goodsImpl) CreateGoods(ctx context.Context, params *CreateGoodsParams) (*models.Goods, error) {
	model := &models.Goods{
		Meta:   models.Meta{},
		Name:   params.Name,
		Weight: params.Weight,
		Length: params.Length,
		Width:  params.Width,
		Height: params.Height,
		Price:  params.Price,
	}
	model.UpdateMeta()

	id, err := g.svc.create(ctx, tableNameGoods, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

type GetGoodsParams struct {
	UserID *string
	Name   *string
}

func (p *GetGoodsParams) filter(sd *goqu.SelectDataset) (*goqu.SelectDataset, error) {
	switch {
	case p.UserID != nil:
		return sd.Where(goqu.Ex{"id": *p.UserID}), nil
	case p.Name != nil:
		return sd.Where(goqu.Ex{"name": p.Name}), nil
	default:
		return nil, errors.New("undefined behavior: id is not set and name is not set")
	}
}

func (g *goodsImpl) GetGoods(ctx context.Context, params *GetGoodsParams) (*models.Goods, error) {
	sd, err := params.filter(goqu.From(tableNameGoods).Select("*"))
	if err != nil {
		return nil, fmt.Errorf("failed to create filter: %w", err)
	}

	query, _, err := sd.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	result := &models.Goods{}
	if err = g.svc.GetContext(ctx, result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}

		return nil, err
	}

	return result, nil
}

func (g *goodsImpl) DeleteGoods(ctx context.Context, id string) error {
	return g.svc.delete(ctx, tableNameGoods, id)

}

type ListGoodsFilter struct {
	Pagination *models.Pagination
}

func (f *ListGoodsFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilsPagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

func (g *goodsImpl) ListGoods(ctx context.Context, filter *ListGoodsFilter) ([]*models.Goods, error) {
	ds := goqu.From(tableNameGoods).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.Goods, 0)
	if err = g.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}
