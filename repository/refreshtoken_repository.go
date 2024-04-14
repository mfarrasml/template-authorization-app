package repository

import (
	"context"
	"database/sql"

	"github.com/mfarrasml/template-authorization-app/apperror"
)

type RefreshTokenRepository interface {
	FindOneByJtiId(ctx context.Context, jti string) error
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

func (r *RefreshTokenRepoPostgres) FindOneByJtiId(ctx context.Context, jti string) error {
	q := `
		SELECT DISTINCT id FROM refresh_tokens
		WHERE jti=$1 AND deleted_at IS NULL
		ORDER BY created_at DESC;
	`
	var refTokenId string
	err := r.db.QueryRowContext(ctx, q, jti).Scan(&refTokenId)
	if err == sql.ErrNoRows {
		return apperror.ErrRefreshTokenNotFound
	}
	if err != nil {
		return apperror.ErrInternalServer
	}

	return nil
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
