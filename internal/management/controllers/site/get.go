package site

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/models"
)

func (s *ServiceImpl) Get(ctx context.Context, id string) (*models.Site, error) {
	return s.db.Site().GetSite(ctx, id)
}
