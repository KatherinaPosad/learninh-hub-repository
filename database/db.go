package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

//$env:conn_string = "postgres://user:password@host:5432/dbname?sslmode=disable"

func ConnectDB(ctx context.Context) (*pgx.Conn, error) {
	connString := os.Getenv("DB_URL")
	fmt.Println("DB_URL:", connString)
	return pgx.Connect(ctx, connString)
}
