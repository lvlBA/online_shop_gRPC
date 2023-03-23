package region

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListRegion(ctx context.Context, params *ListParams) ([]*models.Region, error) {
	resp, err := s.db.Region().ListRegion(ctx, &db.ListRegionFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
