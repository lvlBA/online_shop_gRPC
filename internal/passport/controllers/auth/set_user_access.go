package auth

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) SetUserAccess(ctx context.Context, resourceID string, UserID string) error {
	return controllers.AdaptingErrorDB(s.db.Auth().SetUserAccess(ctx, resourceID, UserID))
}
