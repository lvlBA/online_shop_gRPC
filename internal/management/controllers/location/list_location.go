package location

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListLocation(ctx context.Context, params *ListParams) ([]*models.Location, error) {
	resp, err := s.db.Location().ListLocation(ctx, &db.ListLocationFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
