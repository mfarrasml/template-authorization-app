package repository

import (
	"context"
	"database/sql"

	"github.com/mfarrasml/template-authorization-app/apperror"
)

type RefreshTokenRepository interface {
	FindOneById(ctx context.Context, id int) (*string, error)
}

type RefreshTokenRepoPostgres struct {
	db *sql.DB
}

func NewRefreshTokenRepoPostgres(db *sql.DB) *RefreshTokenRepoPostgres {
	return &RefreshTokenRepoPostgres{
		db: db,
	}
}

func (r *RefreshTokenRepoPostgres) FindOneById(ctx context.Context, id int) (*string, error) {
	q := `
		SELECT token FROM refresh_tokens
		WHERE id=$1;
	`
	var refreshToken string
	err := r.db.QueryRowContext(ctx, q, id).Scan(&refreshToken)
	if err == sql.ErrNoRows {
		return nil, apperror.ErrRefreshTokenNotFound
	}
	if err != nil {
		return nil, apperror.ErrInternalServer
	}

	return &refreshToken, nil
}
