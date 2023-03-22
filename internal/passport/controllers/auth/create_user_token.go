package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type CreateUserTokenRequest struct {
	UserID string
}

func (s *ServiceImpl) CreateUserToken(ctx context.Context, params *CreateUserTokenRequest) (*models.Auth, error) {
	auth, err := s.db.Auth().GetUserAuth(ctx, &db.GetUserAuthParams{
		UserID: &params.UserID,
	})
	if err != nil {
		if !errors.Is(err, db.ErrorNotFound) {
			return nil, fmt.Errorf("failed to get auth by user id: %w", err)
		}

		auth, err = s.db.Auth().CreateUserAuth(ctx, &db.CreateUserTokenParams{
			UserID: params.UserID,
			Token:  createToken(),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create auth: %w", err)
		}
	}

	return auth, nil
}

func createToken() string {
	return uuid.New().String()
}
