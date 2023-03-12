package user

import (
	"context"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

type Service interface {
	CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error)
	GetUser(ctx context.Context, params *GetUserParams) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, params *ListParams) ([]*models.User, error)
	ChangePass(ctx context.Context, id string, oldPass string, newPass string) error
}
