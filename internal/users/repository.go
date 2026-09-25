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

	query := `
	  			INSERT INTO users (username, email, password_hash) 
				VALUES ($1, $2, $3) 
				RETURNING id, username, email, created_at
			`

	err := r.db.QueryRow(ctx, query, username, email, passwordHash).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.CreatedAt,
	)

	return user, err
}

func (repo *Repository) GetByEmail(
	ctx context.Context,
	email string,
) (User, string, error) {

	var user User
	var passwordHash string

	query := `
        SELECT id, username, email, password_hash, created_at
        FROM users
        WHERE email = $1
    `

	err := repo.db.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&passwordHash,
		&user.CreatedAt,
	)

	return user, passwordHash, err
}
