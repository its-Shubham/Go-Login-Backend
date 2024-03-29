package database

import (
	"backend/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*sql.DB
}

// NewDatabase initializes and returns a new PostgreSQL database connection.
func NewDatabase(cfg *config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	//connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}
