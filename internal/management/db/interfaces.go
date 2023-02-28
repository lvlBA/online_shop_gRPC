package db

import (
	"context"

	sql "github.com/jmoiron/sqlx"
	"github.com/lvlBA/online_shop/internal/management/models"
)

type SqlClient interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type Site interface {
	CreateSite(ctx context.Context, params *CreateSiteParams) (*models.Site, error)
	GetSite(ctx context.Context, id string) (*models.Site, error)
	DeleteSite(ctx context.Context, id string) error
	ListSites(ctx context.Context, filter *ListSitesFilter) ([]*models.Site, error)
}

type Service interface {
	Site() Site
}

type service interface {
	Service
	create(ctx context.Context, table string, req any) (string, error)
	update(ctx context.Context, table, id string, req any) error
	delete(ctx context.Context, table, id string) error
}
