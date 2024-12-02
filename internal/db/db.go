package db

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	createRoleQuery = `
		DO $$
		BEGIN
		    CREATE TYPE IF NOT EXISTS role AS ENUM ('admin', 'patient', 'doctor');
		EXCEPTION
		    WHEN duplicate_object THEN
		        -- Type already exists, do nothing
		END $$;
	`
)

func NewDb() (*sql.DB, error) {
	log.Println(os.Getenv("DB_URL"))

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)

	tx, err := db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := tx.Exec(createRoleQuery); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return db, err
}
