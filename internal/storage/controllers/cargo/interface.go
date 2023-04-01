package cargo

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/models"
)

type Service interface {
	CreateCarrier(ctx context.Context, params *CreateParams) (*models.Carrier, error)
	GetCarrier(ctx context.Context, params *GetCargoParams) (*models.Carrier, error)
	DeleteCarrier(ctx context.Context, id string) error
	ListCarrier(ctx context.Context, params *ListParams) ([]*models.Carrier, error)
}
