package region

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) DeleteRegion(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Region().DeleteRegion(ctx, id))
}
