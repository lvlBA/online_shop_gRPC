package orders_store

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type CreateParams struct {
	Name        string
	SiteId      string
	RegionId    string
	LocationId  string
	WarehouseId string
}

func (s *ServiceImpl) CreateOrderStore(ctx context.Context, params *CreateParams) (*models.OrdersStore, error) {
	resp, err := s.db.OrdersStore().CreateOrderStore(ctx, &db.CreateOrdersStoreParams{
		Name:        params.Name,
		SiteId:      params.SiteId,
		RegionId:    params.RegionId,
		LocationId:  params.LocationId,
		WarehouseId: params.WarehouseId,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
