package region

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type CreateParams struct {
	Name string
}

func (s *ServiceImpl) CreateRegion(ctx context.Context, params *CreateParams) (*models.Region, error) {
	resp, err := s.db.Region().CreateRegion(ctx, &db.CreateRegionParams{
		Name: params.Name,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
