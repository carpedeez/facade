package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/carpedeez/facade/config"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func OpenPostgres(ctx context.Context, c config.DBConfig) (Querier, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.Username, c.Password, c.Database)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database connection: %v", err)
	}

	q := postgresQuerier{DB: db}
	return q, nil
}

func (q postgresQuerier) Close() error {
	return q.DB.Close()
}
