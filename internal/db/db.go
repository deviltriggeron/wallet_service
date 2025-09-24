package db

import (
	"database/sql"
	"fmt"
	e "wallet/internal/entity"

	_ "github.com/lib/pq"
)

func Connect(cfg *e.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
	)
	return sql.Open("postgres", dsn)
}
