package user

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"
)

func (s *ServiceImpl) ChangePass(ctx context.Context, id string, oldPass string, newPass string) {
	return controllers.AdaptingErrorDB(s.db.User().ChangePass(ctx, id, oldPass, newPass))
}
