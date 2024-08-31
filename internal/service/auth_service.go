package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitsingh/forky/backend/internal/model"
	"github.com/nikitsingh/forky/backend/internal/repo"
	"golang.org/x/exp/rand"
)

type AuthService struct {
	repo *repo.AuthRepo
}

func NewAuthService(repo *repo.AuthRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateMagicLink(ctx context.Context, email string) (model.MagicLink, error) {
	err := s.repo.MarkPreviousMagicLinksAsUsed(ctx, email)
	if err != nil {
		return model.MagicLink{}, err
	}

	otp := fmt.Sprintf("%06d", rand.Intn(1000000))
	magicLink, err := s.repo.CreateMagicLink(ctx, email, otp)
	if err != nil {
		return model.MagicLink{}, err
	}

	return magicLink, nil
}

func (s *AuthService) VerifyMagicLink(ctx context.Context, email, otp string) error {
	_, err := s.repo.GetMagicLink(ctx, email, otp)
	if err != nil {
		return err
	}

	err = s.repo.MarkMagicLinkAsUsed(ctx, email, otp)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) CreateSession(ctx context.Context, userID uuid.UUID) (model.Session, error) {
	session, err := s.repo.CreateSession(ctx, userID)
	if err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *AuthService) DeleteSession(ctx context.Context, token uuid.UUID) error {
	err := s.repo.DeleteSession(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
