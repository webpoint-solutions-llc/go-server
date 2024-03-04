package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var PgPool *pgxpool.Pool

func New(ctx context.Context) *pgxpool.Pool {
	pgUrl := config.Env.PGURL

	pool, err := pgxpool.New(ctx, pgUrl)
	if err != nil {
		log.Fatalln("database error: %w ", err)
		return nil
	}

	pool.Config().MaxConns = 10
	pool.Config().MaxConnIdleTime = 20 * time.Second

	PgPool = pool

	return pool
}
