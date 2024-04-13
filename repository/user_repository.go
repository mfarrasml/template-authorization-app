package repository

import (
	"context"
	"database/sql"

	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/entity"
)

type UserRepository interface {
	FindOneByEmail(ctx context.Context, email string) (*entity.User, error)
	FindOneById(ctx context.Context, id int) (*entity.User, error)
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
		return nil, apperror.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepoPostgres) FindOneById(ctx context.Context, id int) (*entity.User, error) {
	q := `
		SELECT user_name, email, password, created_at, updated_at FROM users
		WHERE id = $1
	`

	user := entity.User{
		Id: id,
	}

	err := r.db.QueryRowContext(ctx, q, id).Scan(&user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, apperror.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
