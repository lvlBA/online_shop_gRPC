package cargo

import (
	"context"

	"github.com/lvlBA/online_shop/internal/storage/controllers"
	"github.com/lvlBA/online_shop/internal/storage/db"
	"github.com/lvlBA/online_shop/internal/storage/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListCarrier(ctx context.Context, params *ListParams) ([]*models.Carrier, error) {
	resp, err := s.db.Cargo().ListCarrier(ctx, &db.ListCargoFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
