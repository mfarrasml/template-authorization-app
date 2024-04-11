package entity

import "database/sql"

type User struct {
	Id        int
	Name      string
	Email     string
	Password  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
