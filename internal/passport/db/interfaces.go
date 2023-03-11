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
	ListUsers(ctx context.Context, filter *ListUserFilter) ([]*models.User, error)
	ChangePass(ctx context.Context, id string, oldPass string, newPass string) error
}

type Service interface {
	User() User
	Resource() Resource
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

type Resource interface {
	CreateResource(ctx context.Context, params *CreateServiceParams) (*models.Resource, error)
	GetResource(ctx context.Context, id string) (*models.Resource, error)
	DeleteResource(ctx context.Context, id string) error
	ListResource(ctx context.Context, filter *ListServiceFilter) ([]*models.Resource, error)
	SetUserAccess(ctx context.Context, id string) error
	DeleteUserAccess(ctx context.Context, id string) error
}
