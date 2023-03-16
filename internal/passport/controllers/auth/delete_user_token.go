package auth

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) DeleteUserToken(ctx context.Context, userId string) (err error) {
	return controllers.AdaptingErrorDB(s.db.Auth().DeleteUserAuth(ctx, userId))
}
