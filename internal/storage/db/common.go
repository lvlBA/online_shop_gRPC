package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/doug-martin/goqu/v9"
)

// serviceImpl - реализует расширение сервиса
type serviceImpl struct {
	sqlClient
}

func (s *serviceImpl) Goods() Goods {
	return &goodsImpl{
		svc: s,
	}
}

func (s *serviceImpl) create(ctx context.Context, table string, req any) (string, error) {
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
		if strings.Contains(err.Error(), "23505") {
			return "", ErrorAlreadyExists
		}
		return "", err
	}

	return id, nil
}

func (s *serviceImpl) update(ctx context.Context, table, id string, req any) error {
	panic("unimplemented")
}

func (s *serviceImpl) delete(ctx context.Context, table, id string) error {
	query, _, err := goqu.From(table).Delete().Where(goqu.Ex{"id": id}).ToSQL()
	if err != nil {
		return err
	}

	res, err := s.ExecContext(ctx, query)
	if err != nil {
		return err
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return ErrorNotFound
	}

	return nil
}
