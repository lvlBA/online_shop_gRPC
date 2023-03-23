package orders_store

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) DeleteOrderStore(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.OrdersStore().DeleteOrderStore(ctx, id))
}
