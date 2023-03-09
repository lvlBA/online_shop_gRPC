package user

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

func (s *ServiceImpl) GetUser(ctx context.Context, id string) (*models.User, error) {
	resp, err := s.db.User().GetUser(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
