package auth

import (
	"context"
	"errors"
	"github.com/lvlBA/online_shop/internal/passport/models"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"
)

func (s *ServiceImpl) UpdateAuth(ctx context.Context, auth *models.Auth) error {
	if err := s.db.Auth().UpdateAuth(ctx, auth); err != nil {
		if errors.Is(err, db.ErrorNotFound) {
			return nil
		}

		return controllers.AdaptingErrorDB(err)
	}

	return nil
}
