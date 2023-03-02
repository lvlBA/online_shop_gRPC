package site

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/db"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type CreateParams struct {
	Name string
}

func (s *ServiceImpl) Create(ctx context.Context, params *CreateParams) (*models.Site, error) {
	return s.db.Site().CreateSite(ctx, &db.CreateSiteParams{
		Name: params.Name,
	})
}
