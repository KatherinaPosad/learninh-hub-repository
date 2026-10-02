package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

//$env:conn_string = "postgres://user:password@host:5432/dbname?sslmode=disable"

func ConnectDB(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("DB_URL")
	return pgx.Connect(ctx, connString)
}
