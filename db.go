package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getConnString() string {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
}

func NewDB() (*pgxpool.Pool, func()) {
	var schema = `
		CREATE TABLE IF NOT EXISTS employees (
			id serial primary key,
			name text,
			position text,
			salary float
		);`
	dsn := getConnString()
	log.Println("connecting to postgres, dsn - ", dsn)

	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Unable to create connection pool, ", err.Error())
	}
	if err := dbpool.Ping(ctx); err != nil {
		log.Fatal("ping failed, err", err.Error())
	}
	_, err = dbpool.Exec(ctx, schema)
	if err != nil {
		log.Fatal("Unable to run schema sql, ", err.Error())
	}

	return dbpool, dbpool.Close
}
