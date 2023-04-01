package db

import (
	"github.com/jmoiron/sqlx"
)

// txImpl - реализует основной сервис с используемой транзакции
type txImpl struct {
	*sqlx.Tx
	*serviceImpl
}
