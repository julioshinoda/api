package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
)

func GetPGConnection() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
