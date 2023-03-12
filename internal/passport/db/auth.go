package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

const tableNameAuth = "auth"

type AuthImplementation struct {
	svc service
}

type CreateUserTokenParams struct {
	UserID string
	Token  string
}

func (a *AuthImplementation) CreateUserAuth(ctx context.Context, params *CreateUserTokenParams) (*models.Auth, error) {
	model := &models.Auth{
		Meta:   models.Meta{},
		UserID: params.UserID,
		Token:  params.Token,
	}
	model.UpdateMeta()

	id, err := a.svc.create(ctx, tableNameAuth, model)
	if err != nil {
		return nil, err
	}
	model.ID = id

	return model, nil
}

type GetUserAuthParams struct {
	UserID *string
	Token  *string
}

func (p *GetUserAuthParams) filter(sd *goqu.SelectDataset) (*goqu.SelectDataset, error) {
	switch {
	case p.UserID != nil:
		return sd.Where(goqu.Ex{"user_id": *p.UserID}), nil
	case p.Token != nil:
		return sd.Where(goqu.Ex{"token": *p.Token}), nil
	default:
		return nil, errors.New("undefined behavior: user id is not set and token is not ser")
	}
}

func (a *AuthImplementation) GetUserAuth(ctx context.Context, params *GetUserAuthParams) (*models.Auth, error) {
	sd, err := params.filter(goqu.From(tableNameAuth).Select("*"))
	if err != nil {
		return nil, fmt.Errorf("failed to calc filter: %w", err)
	}

	query, _, err := sd.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	result := &models.Auth{}
	if err = a.svc.GetContext(ctx, result, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}

		return nil, err
	}

	return result, nil
}

func (a *AuthImplementation) DeleteUserAuth(ctx context.Context, token string) error {
	query, _, err := goqu.From(tableNameAuth).Delete().Where(goqu.Ex{"token": token}).ToSQL()
	if err != nil {
		return err
	}

	res, err := a.svc.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return ErrorNotFound
	}

	return nil
}

func (a *AuthImplementation) DeleteUserToken(ctx context.Context, token string) error {
	return a.svc.update(ctx, tableNameAuth, token, "")
}
