package site

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) List(ctx context.Context, params *ListParams) ([]*models.Site, error) {
	resp, err := s.db.Site().ListSites(ctx, &db.ListSitesFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
