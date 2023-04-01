package goods

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	"github.com/lvlBA/online_shop/internal/storage/db"
	"github.com/lvlBA/online_shop/internal/storage/models"
)

type GetGoodsParams struct {
	ID   *string
	Name *string
}

func (s *ServiceImpl) GetGoods(ctx context.Context, params *GetGoodsParams) (*models.Goods, error) {
	goods, err := s.db.Goods().GetGoods(ctx, &db.GetGoodsParams{
		UserID: params.ID,
		Name:   params.Name,
	})
	if err != nil {
		err = controllers.AdaptingErrorDB(err)
		return nil, err
	}

	return goods, nil
}
