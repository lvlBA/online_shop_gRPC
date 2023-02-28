package site

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/models"
)

type Service interface {
	Create(ctx context.Context, params *CreateParams) (*models.Site, error)
	Get(ctx context.Context, id string) (*models.Site, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, params *ListParams) ([]*models.Site, error)
}
