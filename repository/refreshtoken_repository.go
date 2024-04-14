package repository

import (
	"context"
	"database/sql"

	"github.com/mfarrasml/template-authorization-app/apperror"
	"github.com/mfarrasml/template-authorization-app/entity"
)

type RefreshTokenRepository interface {
	FindOneByUserId(ctx context.Context, userId int) (*entity.RefreshToken, error)
	CreateOne(ctx context.Context, jti string, userId int) error
}

type RefreshTokenRepoPostgres struct {
	db *sql.DB
}

func NewRefreshTokenRepoPostgres(db *sql.DB) *RefreshTokenRepoPostgres {
	return &RefreshTokenRepoPostgres{
		db: db,
	}
}

func (r *RefreshTokenRepoPostgres) FindOneByUserId(ctx context.Context, userId int) (*entity.RefreshToken, error) {
	q := `
		SELECT id, jti FROM refresh_tokens
		WHERE user_id=$1 AND deleted_at IS NULL
		ORDER BY created_at DESC
		LIMIT 1;
	`
	refToken := entity.RefreshToken{
		UserId: userId,
	}
	err := r.db.QueryRowContext(ctx, q, userId).Scan(&refToken.Id, &refToken.Jti)
	if err == sql.ErrNoRows {
		return nil, apperror.ErrRefreshTokenNotFound
	}
	if err != nil {
		return nil, err
	}

	return &refToken, nil
}

func (r *RefreshTokenRepoPostgres) CreateOne(ctx context.Context, jti string, userId int) error {
	q := `
		INSERT INTO refresh_tokens(jti, user_id)
		VALUES ($1, $2);
	`

	_, err := r.db.ExecContext(ctx, q, jti, userId)
	if err != nil {
		return apperror.ErrInternalServer
	}

	return nil
}
