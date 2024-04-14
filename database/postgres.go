package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDb(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dbUrl)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
