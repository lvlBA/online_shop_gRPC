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

func (s *ServiceImpl) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

type ListParams struct {
}

func (s *ServiceImpl) List(ctx context.Context, params *ListParams) ([]*models.Site, error) {
	panic("unimplemented")
}
