package resource

import (
	"context"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type Service interface {
	CreateResource(ctx context.Context, params *CreateServiceParams) (*models.Resource, error)
	GetResource(ctx context.Context, id string) (*models.Resource, error)
	DeleteResource(ctx context.Context, id string) error
	ListResource(ctx context.Context, params *ListParams) ([]*models.Resource, error)
	SetUserAccess(ctx context.Context, id string) error
	DeleteUserAccess(ctx context.Context, id string) error
}
