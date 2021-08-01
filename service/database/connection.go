package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Open - gets PG DB connection
func Open(ctx context.Context, cfg Config) (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)

	connConf, err := pgx.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("pgx.ParseConfig: %w", err)
	}

	connConf.Logger = &pgxLogAdapter{}

	conn, err := pgx.ConnectConfig(ctx, connConf)
	if err != nil {
		return nil, fmt.Errorf("pgx.ConnectConfig: %w", err)
	}

	return conn, nil
}
