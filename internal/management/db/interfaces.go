package db

import (
	"context"
	"database/sql"

	"github.com/lvlBA/online_shop/internal/management/models"
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
	Site() Site
	Region() Region
	Location() Location
	Warehouse() Warehouse
	OrdersStore() OrdersStore
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

// Site - реализация работы с бд для сайта
type Site interface {
	CreateSite(ctx context.Context, params *CreateSiteParams) (*models.Site, error)
	GetSite(ctx context.Context, id string) (*models.Site, error)
	DeleteSite(ctx context.Context, id string) error
	ListSites(ctx context.Context, filter *ListSitesFilter) ([]*models.Site, error)
}

// Region - реализация работы с бд для региона
type Region interface {
	CreateRegion(ctx context.Context, params *CreateRegionParams) (*models.Region, error)
	GetRegion(ctx context.Context, id string) (*models.Region, error)
	DeleteRegion(ctx context.Context, id string) error
	ListRegion(ctx context.Context, filter *ListRegionFilter) ([]*models.Region, error)
}

// Location - реализация работы с бд для локации
type Location interface {
	CreateLocation(ctx context.Context, params *CreateLocationParams) (*models.Location, error)
	GetLocation(ctx context.Context, id string) (*models.Location, error)
	DeleteLocation(ctx context.Context, id string) error
	ListLocation(ctx context.Context, filter *ListLocationFilter) ([]*models.Location, error)
}

// Warehouse - реализация работы с бд для склада
type Warehouse interface {
	CreateWarehouse(ctx context.Context, params *CreateWarehouseParams) (*models.Warehouse, error)
	GetWarehouse(ctx context.Context, id string) (*models.Warehouse, error)
	DeleteWarehouse(ctx context.Context, id string) error
	ListWarehouses(ctx context.Context, filter *ListWareHouseFilter) ([]*models.Warehouse, error)
}

// OrdersStore - реализация работы с бд для пункта выдачи заказов
type OrdersStore interface {
	CreateOrderStore(ctx context.Context, params *CreateOrdersStoreParams) (*models.OrdersStore, error)
	GetOrderStore(ctx context.Context, id string) (*models.OrdersStore, error)
	DeleteOrderStore(ctx context.Context, id string) error
	ListOrderStores(ctx context.Context, filter *ListOrdersStoreFilter) ([]*models.OrdersStore, error)
}
