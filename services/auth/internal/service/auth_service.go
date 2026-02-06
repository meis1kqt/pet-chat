package service

import (
	"chat-app/internal/models"
	"chat-app/pkg/jwt"
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
	jwtSecret string
}

func NewAuthService(log slog.Logger, storage Storage, TokenTTL time.Duration, jwtSecret string) *AuthService {
	return &AuthService{
		log:      &log,
		storage:  storage,
		tokenTTL: TokenTTL,
		jwtSecret: jwtSecret,
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


func (a *AuthService) Login(ctx context.Context, email string, password string) (string, error) {

	a.log.Info("staring login user")

	user, err := a.storage.GetUser(ctx, email)

	if err != nil {
		a.log.Error("get user error", "error", err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(user.PassHash, []byte(password)); err != nil {
		a.log.Error("password","error", err)
		return "", err
	}

	token, err := jwt.NewToken(user, a.tokenTTL, a.jwtSecret)

	if err != nil {
		a.log.Error("token error", "error", err)
		return "", err
	}

	return token, nil
}