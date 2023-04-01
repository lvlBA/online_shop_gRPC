package goods

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	"github.com/lvlBA/online_shop/internal/storage/db"
	"github.com/lvlBA/online_shop/internal/storage/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListGoods(ctx context.Context, params *ListParams) ([]*models.Goods, error) {
	resp, err := s.db.Goods().ListGoods(ctx, &db.ListGoodsFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
