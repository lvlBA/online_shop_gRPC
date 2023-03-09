package models

type User struct {
	Meta
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	Age          uint32 `db:"age"`
	Sex          Sex    `db:"sex"`
	Login        string `db:"login"`
	HashPassword string `db:"hash_password"`
}
