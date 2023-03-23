package warehouse

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type CreateParams struct {
	Name string
}

func (s *ServiceImpl) CreateWarehouse(ctx context.Context, params *CreateParams) (*models.Warehouse, error) {

	resp, err := s.db.Warehouse().CreateWarehouse(ctx, &db.CreateWarehouseParams{
		Name: params.Name,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
