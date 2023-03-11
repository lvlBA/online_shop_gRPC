package resource

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

func (s *ServiceImpl) GetResource(ctx context.Context, id string) (*models.Resource, error) {
	resp, err := s.db.Resource().GetResource(ctx, id)
	return resp, controllers.AdaptingErrorDB(err)
}
