package orders_store

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/models"
)

func (s *ServiceImpl) GetOrderStore(ctx context.Context, id string) (*models.OrdersStore, error) {
	resp, err := s.db.OrdersStore().GetOrderStore(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
