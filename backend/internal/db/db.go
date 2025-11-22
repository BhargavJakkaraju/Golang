package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func Connect() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not in env")
	}

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Error opening DB", err)
	}

	if err = conn.PingContext(context.Background()); err != nil {
		log.Fatal("Could not ping db: ", err)
	}

	fmt.Println("Connected to PostgreSQL")
	DB = conn
	return DB
}