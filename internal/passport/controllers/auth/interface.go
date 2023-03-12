package auth

import (
	"context"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

type Service interface {
	CreateUserToken(ctx context.Context, params *CreateUserTokenRequest) (*models.Auth, error)
	GetUserToken(ctx context.Context, params *GetUserTokenRequest) (*models.Auth, error)
	DeleteUserToken(ctx context.Context, token string) (err error)
	CheckUserAccess(ctx context.Context, params *CheckUserAccessRequest) (bool, error)
}
