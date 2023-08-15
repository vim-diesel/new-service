package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vim-diesel/new-service/business/data/dbmigrate"
	database "github.com/vim-diesel/new-service/business/sys/database/pgx"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	// Load in the `.env` file
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("loading .env: %w", err)
	}

	dsn := os.Getenv("DSN")

	if err := migrate(dsn); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}

	if err := seed(dsn); err != nil {
		return fmt.Errorf("seed: %w", err)
	}

	return nil
}

func migrate(dsn string) error {
	db, err := database.Open(dsn)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbmigrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}

func seed(dsn string) error {
	db, err := database.Open(dsn)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := dbmigrate.Seed(ctx, db); err != nil {
		return fmt.Errorf("seed database: %w", err)
	}

	fmt.Println("seed data complete")
	return nil
}
