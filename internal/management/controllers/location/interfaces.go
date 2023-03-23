package location

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type Service interface {
	CreateLocation(ctx context.Context, params *CreateParams) (*models.Location, error)
	GetLocation(ctx context.Context, id string) (*models.Location, error)
	DeleteLocation(ctx context.Context, id string) error
	ListLocation(ctx context.Context, params *ListParams) ([]*models.Location, error)
}
