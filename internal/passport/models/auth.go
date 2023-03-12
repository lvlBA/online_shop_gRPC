package models

type Auth struct {
	Meta
	UserID string `db:"user_id"`
	Token  string `db:"token"`
}

type Access struct {
	Meta
	UserID     string `db:"user_id"`
	ResourceID string `db:"resource_id"`
}
