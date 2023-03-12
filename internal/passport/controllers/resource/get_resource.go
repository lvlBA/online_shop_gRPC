package resource

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

func (s *ServiceImpl) GetResourceByID(ctx context.Context, id string) (*models.Resource, error) {
	resp, err := s.db.Resource().GetResource(ctx, &db.GetResourceParams{
		ID: &id,
	})

	return resp, controllers.AdaptingErrorDB(err)
}

func (s *ServiceImpl) GetResourceByName(ctx context.Context, name string) (*models.Resource, error) {
	resp, err := s.db.Resource().GetResource(ctx, &db.GetResourceParams{
		Resource: &name,
	})

	return resp, controllers.AdaptingErrorDB(err)
}
