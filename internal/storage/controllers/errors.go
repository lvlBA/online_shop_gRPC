package controllers

import (
	"errors"

	"github.com/lvlBA/online_shop/internal/storage/db"
)

var (
	ErrorNotFound      = errors.New("not found")
	ErrorAlreadyExists = errors.New("already exists")
)

func AdaptingErrorDB(err error) error {
	switch {
	case err == nil:
		return nil
	case errors.Is(err, db.ErrorAlreadyExists):
		return ErrorAlreadyExists
	case errors.Is(err, db.ErrorNotFound):
		return ErrorNotFound
	default:
		return err
	}
}
