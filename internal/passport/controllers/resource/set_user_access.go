package resource

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) SetUserAccess(ctx context.Context, id string) error {
	return controllers.AdaptingErrorDB(s.db.Resource().SetUserAccess(ctx, id))
}
