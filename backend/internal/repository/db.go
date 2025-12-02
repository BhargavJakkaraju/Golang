package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func TestConnection(db *sql.DB) error {
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return fmt.Errorf("failed to get database version: %w", err)
	}

	fmt.Println("Database connection successful!")
	fmt.Println("Database version:", version)
	return nil
}