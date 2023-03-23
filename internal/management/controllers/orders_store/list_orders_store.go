package orders_store

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListOrderStores(ctx context.Context, params *ListParams) ([]*models.OrdersStore, error) {
	resp, err := s.db.OrdersStore().ListOrderStores(ctx, &db.ListOrdersStoreFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
