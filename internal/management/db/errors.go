package db

import "errors"

var (
	ErrorNotFound      = errors.New("not found")
	ErrorAlreadyExists = errors.New("already exists")
)
