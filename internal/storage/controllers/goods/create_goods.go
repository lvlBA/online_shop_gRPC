package goods

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	"github.com/lvlBA/online_shop/internal/storage/db"
	"github.com/lvlBA/online_shop/internal/storage/models"
)

type CreateParams struct {
	Name   string
	Weight int
	Length int
	Width  int
	Height int
	Price  float64
}

func (s *ServiceImpl) CreateGoods(ctx context.Context, params *CreateParams) (*models.Goods, error) {
	resp, err := s.db.Goods().CreateGoods(ctx, &db.CreateGoodsParams{
		Name:   params.Name,
		Weight: params.Weight,
		Length: params.Length,
		Width:  params.Width,
		Height: params.Height,
		Price:  params.Price,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
