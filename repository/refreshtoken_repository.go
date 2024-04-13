package repository

import (
	"context"
	"database/sql"

	"github.com/mfarrasml/template-authorization-app/apperror"
)

type RefreshTokenRepository interface {
	FindOneByJtiId(ctx context.Context, jti string) (*string, error)
	CreateOne(ctx context.Context, jti string) error
}

type RefreshTokenRepoPostgres struct {
	db *sql.DB
}

func NewRefreshTokenRepoPostgres(db *sql.DB) *RefreshTokenRepoPostgres {
	return &RefreshTokenRepoPostgres{
		db: db,
	}
}

func (r *RefreshTokenRepoPostgres) FindOneByJtiId(ctx context.Context, jti string) (*string, error) {
	q := `
		SELECT DISTINCT token FROM refresh_tokens
		WHERE jti=$1 AND deleted_at IS NULL
		ORDER BY created_at DESC;
	`
	var refreshToken string
	err := r.db.QueryRowContext(ctx, q, jti).Scan(&refreshToken)
	if err == sql.ErrNoRows {
		return nil, apperror.ErrRefreshTokenNotFound
	}
	if err != nil {
		return nil, apperror.ErrInternalServer
	}

	return &refreshToken, nil
}

func (r *RefreshTokenRepoPostgres) CreateOne(ctx context.Context, jti string) error {
	q := `
		INSERT INTO refresh_tokens(jti)
		VALUES ($1);
	`

	_, err := r.db.ExecContext(ctx, q, jti)
	if err != nil {
		return apperror.ErrInternalServer
	}

	return nil
}
