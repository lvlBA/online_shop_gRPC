package user

import (
	"context"
	"crypto/sha512"

	"github.com/lvlBA/online_shop/internal/passport/db"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type CreateUserParams struct {
	FirstName string
	LastName  string
	Age       uint64
	Sex       models.Sex
	Login     string
	Password  string
}

func (s *ServiceImpl) Create(ctx context.Context, params *CreateUserParams) (*models.User, error) {
	hash := sha512.Sum512([]byte(params.Password))
	hashPass := make([]byte, len(hash))
	copy(hashPass, hash[:len(hash)])

	return s.db.User().CreateUser(ctx, &db.CreateUserParams{
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		Age:          params.Age,
		Sex:          params.Sex,
		Login:        params.Login,
		HashPassword: string(hashPass),
	})
}
