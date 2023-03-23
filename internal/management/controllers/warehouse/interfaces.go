package warehouse

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/models"
)

type Service interface {
	CreateWarehouse(ctx context.Context, params *CreateParams) (*models.Warehouse, error)
	GetWarehouse(ctx context.Context, id string) (*models.Warehouse, error)
	DeleteWarehouse(ctx context.Context, id string) error
	ListWarehouses(ctx context.Context, params *ListParams) ([]*models.Warehouse, error)
}
