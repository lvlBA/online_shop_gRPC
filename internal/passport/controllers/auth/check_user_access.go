package auth

import (
	"context"
	"errors"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"
)

type CheckUserAccessRequest struct {
	UserID     string
	ResourceID string
}

func (s *ServiceImpl) CheckUserAccess(ctx context.Context, params *CheckUserAccessRequest) (bool, error) {
	if _, err := s.db.Auth().GetUserAccess(ctx, &db.GetUserAccessParams{
		UserID:     &params.UserID,
		ResourceID: &params.ResourceID,
	}); err != nil {
		if errors.Is(err, db.ErrorNotFound) {
			return false, nil
		}

		return false, controllers.AdaptingErrorDB(err)
	}

	return true, nil
}
