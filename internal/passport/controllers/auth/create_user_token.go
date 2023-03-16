package auth

import (
	"context"
	"crypto/sha512"
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

		token, err := createToken()
		if err != nil {
			return nil, fmt.Errorf("failed to create token: %w", err)
		}

		auth, err = s.db.Auth().CreateUserAuth(ctx, &db.CreateUserTokenParams{
			UserID: params.UserID,
			Token:  token,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create auth: %w", err)
		}
	}

	return auth, nil
}

func createToken() ([]byte, error) {
	id := uuid.New()
	b, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}

	token := sha512.Sum512(b)

	return token[:], nil
}
