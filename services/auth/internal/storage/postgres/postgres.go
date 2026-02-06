package postgres

import (
	"chat-app/services/auth/internal/domain/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)


type Storage struct {
	pool *pgxpool.Pool
}


func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{pool: pool}
}


func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {

	var id int64

	query := "INSERT INTO users (email, passHash) VALUES ($1, $2) returning id"

	err := s.pool.QueryRow(ctx, query, email, passHash).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to save user: %e", err)
	}

	return id, nil
}


func (s *Storage) GetUser(ctx context.Context, email string) (*models.User, error) {
	var User models.User

	query := "SELECT id, email, passHash from users WHERE email = $1"

	if err := s.pool.QueryRow(ctx, query, email).Scan(&User.Email, &User.PassHash, &User.ID); err != nil {
		return &models.User{}, err
	}

	return &User, nil
}

