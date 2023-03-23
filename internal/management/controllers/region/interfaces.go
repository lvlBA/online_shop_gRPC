package region

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type Service interface {
	CreateRegion(ctx context.Context, params *CreateParams) (*models.Region, error)
	GetRegion(ctx context.Context, id string) (*models.Region, error)
	DeleteRegion(ctx context.Context, id string) error
	ListRegion(ctx context.Context, params *ListParams) ([]*models.Region, error)
}
