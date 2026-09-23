package internal

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {

	database_url := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), database_url)

	if err != nil {
		return nil, err
	}

	if err = pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
