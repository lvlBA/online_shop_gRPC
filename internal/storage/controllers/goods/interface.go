package goods

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/models"
)

type Service interface {
	CreateGoods(ctx context.Context, params *CreateParams) (*models.Goods, error)
	GetGoods(ctx context.Context, params *GetGoodsParams) (*models.Goods, error)
	DeleteGoods(ctx context.Context, id string) error
	ListGoods(ctx context.Context, params *ListParams) ([]*models.Goods, error)
}
