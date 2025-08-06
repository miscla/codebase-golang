package database

import (
	"fmt"

	"codebase-golang/pkg/config"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func Init(cfg config.Config) error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return err
	}

	DB = db
	return nil
}
