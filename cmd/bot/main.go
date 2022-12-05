package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func main() {
	dsn := "postgresql://bot_dev:s1mOH8z9FkSK3i4abAkOiA@gamma-lemur-6172.8nj.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dsn)
	defer conn.Close()
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	var now time.Time
	err = conn.QueryRow(ctx, "SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}
