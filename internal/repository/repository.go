package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewRepository(cfg *Config) (*Repository, error) {
	if cfg.Host == "" {
		return nil, errors.New("host is required")
	}
	if cfg.User == "" {
		return nil, errors.New("user is required")
	}
	if cfg.DBName == "" {
		return nil, errors.New("dbname is required")
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{db: db}, nil
}

func (r *Repository) Close() error {
	if r.db == nil {
		return nil
	}
	return r.db.Close()
}
