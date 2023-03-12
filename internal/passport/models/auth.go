package models

type Auth struct {
	Meta
	UserID string `db:"user_id"`
	Token  string `db:"token"`
}
