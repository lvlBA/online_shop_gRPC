package db

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// ServiceImpl - Основной сервис
type ServiceImpl struct {
	*serviceImpl          // расширяем возможности текущего объекта через наследование
	db           *sqlx.DB // тут нужно именно установить свойство, что бы методы sqlx.DB и serviceImpl не пересекались
}

func New(db *sqlx.DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
		serviceImpl: &serviceImpl{
			sqlClient: db,
		},
	}
}

func (s *ServiceImpl) Begin(ctx context.Context) (Transaction, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &txImpl{
		Tx: tx,
		serviceImpl: &serviceImpl{
			sqlClient: tx,
		},
	}, nil
}
