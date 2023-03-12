package user

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type GetUserParams struct {
	ID       *string
	Login    *string
	Password *string
}

func (s *ServiceImpl) GetUser(ctx context.Context, params *GetUserParams) (*models.User, error) {
	var hash *string
	if params.Password != nil {
		h := toHash(*params.Password)
		hash = &h
	}

	resp, err := s.db.User().GetUser(ctx, &db.GetUserParams{
		ID:           params.ID,
		Login:        params.Login,
		HashPassword: hash,
	})

	return resp, controllers.AdaptingErrorDB(err)
}
