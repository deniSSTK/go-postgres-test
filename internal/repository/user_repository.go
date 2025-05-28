package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"go-postgres-test/internal/domain"
)

type userRepo struct {
	conn *pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) domain.UserRepository {
	return &userRepo{conn: conn}
}

func (r *userRepo) CreateUser(user domain.User) error {
	_, err := r.conn.Exec(context.Background(),
		"INSERT INTO users (username, email) VALUES ($1, $2)",
		user.Username,
		user.Email,
	)
	return err
}
