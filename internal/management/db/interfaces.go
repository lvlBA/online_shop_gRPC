package db

import (
	"context"
	"database/sql"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type SqlClient interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type Site interface {
	CreateSite(ctx context.Context, params *CreateSiteParams) (*models.Site, error)
	GetSite(ctx context.Context, id string) (*models.Site, error)
	DeleteSite(ctx context.Context, id string) error
	ListSites(ctx context.Context, filter *ListSitesFilter) ([]*models.Site, error)
}

type Service interface {
	Site() Site
	Region() Region
	Location() Location
	Warehouse() Warehouse
	OrdersStore() OrdersStore
}

type service interface {
	Service
	SqlClient
	create(ctx context.Context, table string, req any) (string, error)
	update(ctx context.Context, table, id string, req any) error
	delete(ctx context.Context, table, id string) error
}
type Region interface {
	CreateRegion(ctx context.Context, params *CreateRegionParams) (*models.Region, error)
	GetRegion(ctx context.Context, id string) (*models.Region, error)
	DeleteRegion(ctx context.Context, id string) error
	ListRegion(ctx context.Context, filter *ListRegionFilter) ([]*models.Region, error)
}

type Location interface {
	CreateLocation(ctx context.Context, params *CreateLocationParams) (*models.Location, error)
	GetLocation(ctx context.Context, id string) (*models.Location, error)
	DeleteLocation(ctx context.Context, id string) error
	ListLocation(ctx context.Context, filter *ListLocationFilter) ([]*models.Location, error)
}

type Warehouse interface {
	CreateWarehouse(ctx context.Context, params *CreateWarehouseParams) (*models.Warehouse, error)
	GetWarehouse(ctx context.Context, id string) (*models.Warehouse, error)
	DeleteWarehouse(ctx context.Context, id string) error
	ListWarehouses(ctx context.Context, filter *ListWareHouseFilter) ([]*models.Warehouse, error)
}

type OrdersStore interface {
	CreateOrderStore(ctx context.Context, params *CreateOrdersStoreParams) (*models.OrdersStore, error)
	GetOrderStore(ctx context.Context, id string) (*models.OrdersStore, error)
	DeleteOrderStore(ctx context.Context, id string) error
	ListOrderStores(ctx context.Context, filter *ListOrdersStoreFilter) ([]*models.OrdersStore, error)
}
