package resource

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"

	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListResource(ctx context.Context, params *ListParams) ([]*models.Resource, error) {
	resp, err := s.db.Resource().ListResource(ctx, &db.ListServiceFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
