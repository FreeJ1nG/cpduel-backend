package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePool(ctx context.Context, dsn string) (pool *pgxpool.Pool) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("Failed to create database pool", err.Error())
	}
	return pool
}

func TestConnection(ctx context.Context, pool *pgxpool.Pool) {
	_, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database ", err.Error())
	}
	fmt.Println("Connected to database ")
}
