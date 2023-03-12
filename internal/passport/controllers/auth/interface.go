package auth

import (
	"context"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

type Service interface {
	GetUserToken(ctx context.Context, params *GetUserTokenRequest) (*models.Auth, error)
	CheckUser(ctx context.Context, id string) (*models.Auth, error)
}
