package site

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"

	"github.com/lvlBA/online_shop/internal/management/models"
)

func (s *ServiceImpl) Get(ctx context.Context, id string) (*models.Site, error) {
	resp, err := s.db.Site().GetSite(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
