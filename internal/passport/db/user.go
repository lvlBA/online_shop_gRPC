package db

import (
	"context"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

const tableNameUser = "User"

type UserImpl struct {
	svc service
}

type CreateUserParams struct {
	FirstName    string
	LastName     string
	Age          uint32
	Sex          models.Sex
	Login        string
	HashPassword string
}

func (u *UserImpl) CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error) {
	model := &models.User{
		Meta:         models.Meta{},
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		Age:          params.Age,
		Sex:          params.Sex,
		Login:        params.Login,
		HashPassword: params.HashPassword,
	}
	model.UpdateMeta()

	id, err := u.svc.create(ctx, tableNameUser, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}
