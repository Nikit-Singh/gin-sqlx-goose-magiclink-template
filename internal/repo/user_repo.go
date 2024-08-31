package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nikitsingh/forky/backend/internal/model"
)

type UserRepo struct {
	DB *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, email string) (model.User, error) {
	var user model.User
	query := `
		insert into users (email) values ($1) returning *
	`

	err := r.DB.GetContext(ctx, &user, query, email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	query := `
		select * from users where email = $1
	`

	err := r.DB.GetContext(ctx, &user, query, email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
