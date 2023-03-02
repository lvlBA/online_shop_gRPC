package site

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

// мы положим в эту структуру запрос или результат?
type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) List(ctx context.Context, params *ListParams) ([]*models.Site, error) {
	return s.db.Site().ListSites(ctx, &db.ListSitesFilter{
		Pagination: params.Pagination,
	})
}
