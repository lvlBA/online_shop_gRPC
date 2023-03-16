package auth

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type GetUserTokenRequest struct {
	UserID *string
	Token  []byte
}

func (s *ServiceImpl) GetUserToken(ctx context.Context, params *GetUserTokenRequest) (*models.Auth, error) {
	auth, err := s.db.Auth().GetUserAuth(ctx, &db.GetUserAuthParams{
		UserID: params.UserID,
		Token:  params.Token,
	})
	if err != nil {
		return nil, controllers.AdaptingErrorDB(err)
	}

	return auth, nil
}
