package warehouse

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) DeleteWarehouse(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Warehouse().DeleteWarehouse(ctx, id))
}
