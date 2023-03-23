package warehouse

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/models"
)

func (s *ServiceImpl) GetWarehouse(ctx context.Context, id string) (*models.Warehouse, error) {
	resp, err := s.db.Warehouse().GetWarehouse(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
