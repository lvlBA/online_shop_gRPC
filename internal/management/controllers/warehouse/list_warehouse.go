package warehouse

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListWarehouses(ctx context.Context, params *ListParams) ([]*models.Warehouse, error) {
	resp, err := s.db.Warehouse().ListWarehouses(ctx, &db.ListWareHouseFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
