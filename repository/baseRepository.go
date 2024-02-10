package repository

import (
	"context"

	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
)

var conn *pgx.Conn = nil
var contextt = context.Background()

func OpenDB() (*pgx.Conn, error) {

	if conn == nil {
		conn, err := pgx.Connect(context.Background(), "postgres://postgres:pass@db:5432/postgres")
		return conn, err
	}

	return conn, nil
}
