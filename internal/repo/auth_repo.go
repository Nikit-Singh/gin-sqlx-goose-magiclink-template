package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/nikitsingh/forky/backend/internal/model"
)

type AuthRepo struct {
	DB *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

func (r *AuthRepo) CreateMagicLink(ctx context.Context, email, otp string) (model.MagicLink, error) {
	var magicLink model.MagicLink
	query := `
		insert into magic_links (email, otp)
		values ($1, $2)
		returning *
	`

	err := r.DB.GetContext(ctx, &magicLink, query, email, otp)
	if err != nil {
		return model.MagicLink{}, err
	}

	return magicLink, nil
}

func (r *AuthRepo) GetMagicLink(ctx context.Context, email, otp string) (model.MagicLink, error) {
	var magicLink model.MagicLink
	query := `
		select * from magic_links
		where email = $1 and otp = $2 and is_used = false and expires_at > now()
	`

	err := r.DB.GetContext(ctx, &magicLink, query, email, otp)
	if err != nil {
		return model.MagicLink{}, err
	}

	return magicLink, nil
}

func (r *AuthRepo) MarkPreviousMagicLinksAsUsed(ctx context.Context, email string) error {
	query := `
		update magic_links
		set is_used = true
		where email = $1 and is_used = false
	`

	_, err := r.DB.ExecContext(ctx, query, email)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepo) MarkMagicLinkAsUsed(ctx context.Context, email, otp string) error {
	query := `
		update magic_links
		set is_used = true
		where email = $1 and otp = $2 and is_used = false
	`

	_, err := r.DB.ExecContext(ctx, query, email, otp)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepo) CreateSession(ctx context.Context, userID uuid.UUID) (model.Session, error) {
	var session model.Session
	query := `
		insert into sessions (user_id)
		values ($1)
		returning *
	`

	err := r.DB.GetContext(ctx, &session, query, userID)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (r *AuthRepo) DeleteSession(ctx context.Context, token uuid.UUID) error {
	query := `
		delete from sessions
		where token = $1
	`

	_, err := r.DB.ExecContext(ctx, query, token)
	if err != nil {
		return err
	}

	return nil
}
