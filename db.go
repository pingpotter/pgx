package main

import (
	"context"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pingpotter/pgx/logpgx"
)

type queryInterface interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	// Commit(ctx context.Context) error
	// Rollback(ctx context.Context) error
}

type queryInterfacex interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type postgresRepository struct {
	qi  queryInterface
	qix queryInterfacex
}

func connectWithRetries(ctx context.Context, url string) (*pgxpool.Pool, error) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	x := &logpgx.Logger{}
	cfg.ConnConfig.Logger = x
	cfg.ConnConfig.LogLevel = 5
	cfg.MaxConnLifetime = 250 * time.Millisecond

	for {
		pool, err := pgxpool.ConnectConfig(ctx, cfg)
		if err != nil {
			// wait
			select {
			case <-ctx.Done(): // cancellation
				return nil, ctx.Err()
			case <-ticker.C:
				continue
			}
		}
		return pool, nil
	}
}
