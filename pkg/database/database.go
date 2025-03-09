package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mrbelka12000/mock_server/pkg/config"
)

const (
	migrationsDir = "migrations/"
)

// Connect ..
func Connect(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s?_foreign_keys=on", cfg.PathToDB)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	err = runMigrations(db)
	if err != nil {
		return nil, fmt.Errorf("run migrations: %w", err)
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	dir, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("read dir: %w", err)
	}

	for _, file := range dir {
		fileBody, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
		if err != nil {
			return fmt.Errorf("read file: %w", err)
		}

		_, err = db.Exec(string(fileBody))
		if err != nil {
			return fmt.Errorf("exec sql: %w", err)
		}
	}

	return nil
}
