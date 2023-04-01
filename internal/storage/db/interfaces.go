package db

import (
	"context"
	"database/sql"

	"github.com/lvlBA/online_shop/internal/storage/models"
)

// sqlClient - клинт базы данных который мы используем
type sqlClient interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

// Сервисы которые мы возвращаем для работы в контроллере
type services interface {
	Goods() Goods
	Cargo() Cargo
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
	update(ctx context.Context, table, id string, req any) error
	delete(ctx context.Context, table, id string) error
}

// Goods - реализация работы с бд для товаров
type Goods interface {
	CreateGoods(ctx context.Context, params *CreateGoodsParams) (*models.Goods, error)
	GetGoods(ctx context.Context, params *GetGoodsParams) (*models.Goods, error)
	DeleteGoods(ctx context.Context, id string) error
	ListGoods(ctx context.Context, filter *ListGoodsFilter) ([]*models.Goods, error)
}

// Cargo - реализация работы с бд для перевозчиков

type Cargo interface {
	CreateCarrier(ctx context.Context, params *CreateCargoParams) (*models.Carrier, error)
	GetCarrier(ctx context.Context, params *GetCargoParams) (*models.Carrier, error)
	DeleteCarrier(ctx context.Context, id string) error
	ListCarrier(ctx context.Context, filter *ListCargoFilter) ([]*models.Carrier, error)
}
