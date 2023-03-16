package resource

import (
	"context"
	"github.com/lvlBA/online_shop/internal/management/controllers"

	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type CreateResourceParams struct {
	Urn string
}

func (s *ServiceImpl) CreateResource(ctx context.Context, params *CreateResourceParams) (*models.Resource, error) {
	resp, err := s.db.Resource().CreateResource(ctx, &db.CreateResourceParams{
		Urn: params.Urn,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
