package user

import (
	"context"

	"github.com/lvlBA/online_shop/internal/management/controllers"
	"github.com/lvlBA/online_shop/internal/passport/db"
)

func (s *ServiceImpl) DeleteUser(ctx context.Context, id string) (err error) {
	tx, rErr := s.db.Begin(ctx)
	if rErr != nil {
		return rErr
	}
	_ = tx.Auth().DeleteUserAuth(ctx, id)
	_ = tx.Auth().DeleteUserAccess(ctx, &db.DeleteUserAccessParams{
		UserID: &id,
	})

	err = controllers.AdaptingErrorDB(tx.User().DeleteUser(ctx, id))
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return rErr
		}
		return err
	}

	if rErr := tx.Commit(); rErr != nil {
		return rErr
	}

	return nil
}
