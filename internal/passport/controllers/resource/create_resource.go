package resource

import (
	"context"
	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type CreateServiceParams struct {
	Urn string
}

func (s *ServiceImpl) CreateResource(ctx context.Context, params *CreateServiceParams) (*models.Resource, error) {
	return s.db.Resource().CreateResource(ctx, &db.CreateServiceParams{
		Urn: params.Urn,
	})
}
