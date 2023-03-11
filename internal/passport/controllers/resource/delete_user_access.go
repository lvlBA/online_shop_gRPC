package resource

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) DeleteUserAccess(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Resource().DeleteUserAccess(ctx, id))
}
