package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/doug-martin/goqu/v9"

	"github.com/lvlBA/online_shop/internal/passport/models"
	utilspagination "github.com/lvlBA/online_shop/pkg/utils/pagination"
)

const tableNameUser = "users"

type UserImpl struct {
	svc service
}

type CreateUserParams struct {
	FirstName    string
	LastName     string
	Age          uint64
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

type GetUserParams struct {
	ID           *string
	Login        *string
	HashPassword *string
}

func (p *GetUserParams) filter(sd *goqu.SelectDataset) (*goqu.SelectDataset, error) {
	if p.HashPassword != nil {
		sd = sd.Where(goqu.Ex{"hash_password": *p.HashPassword})
	}

	switch {
	case p.ID != nil:
		return sd.Where(goqu.Ex{"id": *p.ID}), nil
	case p.Login != nil:
		return sd.Where(goqu.Ex{"login": *p.Login}), nil
	default:
		return nil, errors.New("undefined behavior: id is not set and login is not set")
	}
}

func (u *UserImpl) GetUser(ctx context.Context, params *GetUserParams) (*models.User, error) {
	sd, err := params.filter(goqu.From(tableNameUser).Select("*"))
	if err != nil {
		return nil, err
	}

	query, _, err := sd.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}

	result := &models.User{}
	if err = u.svc.GetContext(ctx, result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, err
	}

	return result, nil
}

func (u *UserImpl) DeleteUser(ctx context.Context, id string) error {
	return u.svc.delete(ctx, tableNameUser, id)
}

type ListUserFilter struct {
	Pagination *models.Pagination
}

func (f *ListUserFilter) Filter(ds *goqu.SelectDataset) *goqu.SelectDataset {
	if f.Pagination != nil {
		utilspagination.NewPagination(f.Pagination.Page, f.Pagination.Limit).DataSet(ds)
	}

	return ds
}

func (u *UserImpl) ListUsers(ctx context.Context, filter *ListUserFilter) ([]*models.User, error) {
	ds := goqu.From(tableNameUser).Select("*")
	ds = filter.Filter(ds)
	query, _, err := ds.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to create query: %w", err)
	}
	result := make([]*models.User, 0)
	if err = u.svc.SelectContext(ctx, &result, query); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserImpl) ChangePass(ctx context.Context, id string, oldPass string, newPass string) error {
	result := &models.User{}
	query, _, err := goqu.From(tableNameUser).Select("*").Where(goqu.Ex{"id": id}, goqu.Ex{"hash_password": oldPass}).ToSQL()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("failed to create query: %w", err)
	}
	if err = u.svc.GetContext(ctx, result, query); err == nil {

		ds := goqu.From(tableNameUser)
		_, _, _ = ds.Where(goqu.C("id").Eq(result.ID)).Update().Set(
			goqu.Record{"hash_password": newPass},
		).ToSQL()
	}
	return err
}
