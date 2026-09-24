package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, username string, email string, passwordHash string) (User, error) {

	var user User

	query := `INSRT INTO (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, username, email, created_at`

	err := r.db.QueryRow(ctx, query, user, email, passwordHash).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)

	return user, err
}
