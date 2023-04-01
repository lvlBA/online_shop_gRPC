package goods

import (
	"context"
	"github.com/lvlBA/online_shop/internal/storage/controllers"
)

func (s *ServiceImpl) DeleteGoods(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Goods().DeleteGoods(ctx, id))

}
