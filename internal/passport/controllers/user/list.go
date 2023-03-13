package user

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type ListParams struct {
	Pagination *models.Pagination
}

func (s *ServiceImpl) ListUsers(ctx context.Context, params *ListParams) ([]*models.User, error) {
	resp, err := s.db.User().ListUsers(ctx, &db.ListUserFilter{
		Pagination: params.Pagination,
	})
	return resp, controllers.AdaptingErrorDB(err)
}
