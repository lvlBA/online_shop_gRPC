package orders_store

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type CreateParams struct {
	Name string
}

func (s *ServiceImpl) CreateOrderStore(ctx context.Context, params *CreateParams) (*models.OrdersStore, error) {
	resp, err := s.db.OrdersStore().CreateOrderStore(ctx, &db.CreateOrdersStoreParams{
		Name: params.Name,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
