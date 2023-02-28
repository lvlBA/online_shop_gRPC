package db

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

// create - реализует создание объекта по запросу и возвращает его ID
func (s *ServiceImpl) create(ctx context.Context, table string, req any) (string, error) {
	query, _, err := goqu.From(table).
		Insert().
		Rows(req).
		Returning("id").
		ToSQL()
	if err != nil {
		return "", fmt.Errorf("failed to create query: %w", err)
	}

	var id string
	if err = s.sqlClient.QueryRowContext(ctx, query).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (s *ServiceImpl) update(ctx context.Context, table, id string, req any) error {
	panic("unimplemented")
}

func (s *ServiceImpl) delete(ctx context.Context, table, id string) error {
	panic("unimplemented")
}
