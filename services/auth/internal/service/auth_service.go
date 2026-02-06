package service

import (
	"chat-app/services/auth/internal/domain/models"
	"context"
	"errors"
	"log/slog"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (int64, error)
	GetUser(ctx context.Context, email string) (*models.User, error)
}

type AuthService struct {
	log      *slog.Logger
	storage  Storage
	tokenTTL time.Duration
}

func NewAuthService(log slog.Logger, storage Storage, TokenTTL time.Duration) *AuthService {
	return &AuthService{
		log:      &log,
		storage:  storage,
		tokenTTL: TokenTTL,
	}
}

func (a *AuthService) RegisterNewUser(ctx context.Context, email string, password string) (int64, error) {

	a.log.Info("registering user...")

	if password == "" {

		a.log.Error("password is empty", "error", "password is emty")
		return 0, errors.New("password is empty")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		a.log.Error("error with passHAsh", "error", err)
		return 0, err
	}

	id, err := a.storage.SaveUser(ctx, email, passHash)

	if err != nil {
		a.log.Error("failed to save user", "error", err)
		return 0, err
	}

	return id, nil
}
