package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mrbelka12000/mock_server/pkg/config"
)

// Connect ..
func Connect(cfg config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.PGURL)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return db, nil
}
