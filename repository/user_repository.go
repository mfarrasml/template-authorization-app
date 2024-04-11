package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mfarrasml/template-authorization-app/entity"
)

type UserRepository interface {
	FindOneByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserRepoPostgres struct {
	db *sql.DB
}

func NewUserRepoPostgres(db *sql.DB) *UserRepoPostgres {
	return &UserRepoPostgres{
		db: db,
	}
}

func (r *UserRepoPostgres) FindOneByEmail(ctx context.Context, email string) (*entity.User, error) {
	q := `
		SELECT id, user_name, email, password FROM users
		WHERE email = $1
	`

	user := entity.User{}

	err := r.db.QueryRowContext(ctx, q, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("error user %s not found", email)
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
