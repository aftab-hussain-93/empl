package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDB() (*pgxpool.Pool, error) {
	// var schema = `
	// 	CREATE TABLE IF NOT EXISTS employee (
	// 		id int,
	// 		name text,
	// 		position text,
	// 		salary float64
	// 	);`

	connString := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
	if dsn := os.Getenv("POSTGRES_URL"); dsn != "" {
		connString = dsn
	}

	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// dbpool.Qu
	// // pgxpool.MustExec(s÷chema)
	return dbpool, nil
}
