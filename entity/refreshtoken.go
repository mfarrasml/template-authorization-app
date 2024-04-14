package entity

import "database/sql"

type RefreshToken struct {
	Id        int
	UserId    int
	Jti       string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
