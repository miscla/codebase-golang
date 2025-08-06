package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"codebase-golang/pkg/config"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	// optional: configure pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	// ping to verify
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("connected to postgres")
	return db, nil
}
