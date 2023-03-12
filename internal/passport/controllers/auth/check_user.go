package auth

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type CheckUserParams struct {
	Id    string
	Login string
}

func (s *ServiceImpl) CheckUser(ctx context.Context, id string) (*models.Auth, error) {
	resp, err := s.db.Auth().CheckUser(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
