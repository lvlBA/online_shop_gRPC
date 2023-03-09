package db

import (
	"context"
	"database/sql"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

type User interface {
	CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	//ListUsers(ctx context.Context, filter *ListUserFilters) ([]*models.User, error)
	ChangePass(ctx context.Context, login string) error
}

type Service interface {
	User() User
}

type SqlClient interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type service interface {
	Service
	SqlClient
	create(ctx context.Context, table string, req any) (string, error)
	update(ctx context.Context, table, id string, req any) error
	delete(ctx context.Context, table, id string) error
}
