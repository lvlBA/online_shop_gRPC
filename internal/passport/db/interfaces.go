package db

import (
	"context"
	"database/sql"

	"github.com/lvlBA/online_shop/internal/passport/models"
)

// sqlClient  - клинт базы данных который мы используем
type sqlClient interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

// Сервисы которые мы возвращаем для работы в контроллере
type services interface {
	User() User
	Resource() Resource
	Auth() Auth
}

// Service - основной интерфейс сервиса
type Service interface {
	services
	Begin(ctx context.Context) (Transaction, error)
}

// Transaction - Основной интерйес сервиса, при работе через транзакцию
type Transaction interface {
	services
	Commit() error
	Rollback() error
}

// sqlService - Расширенная версия sql клиента предоставляемая сервисом
type sqlService interface {
	services
	sqlClient
	create(ctx context.Context, table string, req any) (string, error)
	update(ctx context.Context, table string, id string, req interface{}) error
	delete(ctx context.Context, table, id string) error
}

// User - реализация работы с бд для пользователя
type User interface {
	CreateUser(ctx context.Context, params *CreateUserParams) (*models.User, error)
	GetUser(ctx context.Context, params *GetUserParams) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, filter *ListUserFilter) ([]*models.User, error)
	ChangePass(ctx context.Context, id string, oldPass string, newPass string) error
}

// Resource - реализация работы с бд для ресурса
type Resource interface {
	CreateResource(ctx context.Context, params *CreateResourceParams) (*models.Resource, error)
	GetResource(ctx context.Context, params *GetResourceParams) (*models.Resource, error)
	DeleteResource(ctx context.Context, id string) error
	ListResource(ctx context.Context, filter *ListServiceFilter) ([]*models.Resource, error)
}

// Auth - реализация работы с бд для аутентификации
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
