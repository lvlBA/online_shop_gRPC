package user

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) DeleteUser(ctx context.Context, id string) (err error) {
	return controllers.AdaptingErrorDB(s.db.User().DeleteUser(ctx, id))
}
