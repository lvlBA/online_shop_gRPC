package auth

import (
	"context"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

type Service interface {
	CreateUserToken(ctx context.Context, params *CreateUserTokenRequest) (*models.Auth, error)
	GetUserToken(ctx context.Context, params *GetUserTokenRequest) (*models.Auth, error)
	DeleteUserToken(ctx context.Context, userId string) (err error)
	CheckUserAccess(ctx context.Context, params *CheckUserAccessRequest) (bool, error)
	SetUserAccess(ctx context.Context, resourceID string, UserID string) error
	DeleteUserAccess(ctx context.Context, userID *string, resourceId *string) error
	UpdateAuth(ctx context.Context, auth *models.Auth) error
}
