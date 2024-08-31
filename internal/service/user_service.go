package service

import (
	"context"

	"github.com/nikitsingh/forky/backend/internal/model"
	"github.com/nikitsingh/forky/backend/internal/repo"
)

type UserService struct {
	repo *repo.UserRepo
}

func NewUserService(repo *repo.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, email string) (model.User, error) {
	user, err := s.repo.CreateUser(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
