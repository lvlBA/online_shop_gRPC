package db

import (
	"context"
	"database/sql"
	"github.com/lvlBA/online_shop/internal/passport/models"
)

type User interface {
	CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error)
	GetUser(ctx context.Context, params *GetUserParams) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, filter *ListUserFilter) ([]*models.User, error)
	ChangePass(ctx context.Context, id string, oldPass string, newPass string) error
}

type Service interface {
	User() User
	Resource() Resource
	Auth() Auth
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
	update(ctx context.Context, table string, id string, req interface{}) error
	delete(ctx context.Context, table, id string) error
}

type Resource interface {
	CreateResource(ctx context.Context, params *CreateResourceParams) (*models.Resource, error)
	GetResource(ctx context.Context, params *GetResourceParams) (*models.Resource, error)
	DeleteResource(ctx context.Context, id string) error
	ListResource(ctx context.Context, filter *ListServiceFilter) ([]*models.Resource, error)
}

type Auth interface {
	CreateUserAuth(ctx context.Context, params *CreateUserTokenParams) (*models.Auth, error)
	GetUserAuth(ctx context.Context, params *GetUserAuthParams) (*models.Auth, error)
	DeleteUserAuth(ctx context.Context, userId string) error

	CreateUserAccess(ctx context.Context, params *CreateUserAccessParams) (*models.Access, error)
	DeleteUserAccess(ctx context.Context, params *DeleteUserAccessParams) error
	SetUserAccess(ctx context.Context, resourceID string, UserID string) error
	GetUserAccess(ctx context.Context, params *GetUserAccessParams) (*models.Access, error)
	UpdateAuth(ctx context.Context, auth *models.Auth) error
	DeleteOldTokens(ctx context.Context) error
}
