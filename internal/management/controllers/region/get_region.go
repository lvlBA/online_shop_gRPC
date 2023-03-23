package region

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/models"
)

func (s *ServiceImpl) GetRegion(ctx context.Context, id string) (*models.Region, error) {
	resp, err := s.db.Region().GetRegion(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
