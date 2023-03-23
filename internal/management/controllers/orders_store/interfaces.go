package orders_store

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/models"
)

type Service interface {
	CreateOrderStore(ctx context.Context, params *CreateParams) (*models.OrdersStore, error)
	GetOrderStore(ctx context.Context, id string) (*models.OrdersStore, error)
	DeleteOrderStore(ctx context.Context, id string) error
	ListOrderStores(ctx context.Context, params *ListParams) ([]*models.OrdersStore, error)
}
