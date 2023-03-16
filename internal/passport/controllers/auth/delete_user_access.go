package auth

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"
)

type DeleteUserAccessParams struct {
	UserID     *string
	ResourceID *string
}

func (s *ServiceImpl) DeleteUserAccess(ctx context.Context, userID *string, resourceId *string) error {
	return controllers.AdaptingErrorDB(s.db.Auth().DeleteUserAccess(ctx, &db.DeleteUserAccessParams{
		UserID:     userID,
		ResourceID: resourceId,
	}))
}
