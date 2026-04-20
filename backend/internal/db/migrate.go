package db

import (
	"database/sql"
	"embed"
	"fmt"
	"time"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func RunMigrations(db *sql.DB) error {
	// Create schema_migrations table first if it doesn't exist
	_, err := db.Exec(`
	  CREATE TABLE IF NOT EXISTS schema_migrations (
		  version   TEXT PRIMARY KEY,
		  applied_at TIMESTAMPTZ NOT NULL
	  )
	  `)
	if err != nil {
		return fmt.Errorf("create schema_migrations table: %w", err)
	}

	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		version := entry.Name() // e.g. "001_init.sql"
		if alreadyApplied(db, version) {
			continue
		}

		sqlBytes, err := migrationsFS.ReadFile("migrations/" + version)
		if err != nil {
			return fmt.Errorf("read %s: %w", version, err)
		}

		tx, err := db.Begin()
		if err != nil {
			return fmt.Errorf("begin transaction for %s: %w", version, err)
		}

		if _, err := tx.Exec(string(sqlBytes)); err != nil {
			tx.Rollback()
			return fmt.Errorf("execute %s: %w", version, err)
		}

		if _, err := tx.Exec(
			"INSERT INTO schema_migrations VALUES (?, ?)",
			version, time.Now(),
		); err != nil {
			tx.Rollback()
			return fmt.Errorf("record migration %s: %w", version, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit %s: %w", version, err)
		}

		fmt.Printf("Applied migration: %s\n", version)
	}

	return nil
}

func alreadyApplied(db *sql.DB, version string) bool {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM schema_migrations WHERE version = ?", version).Scan(&count)
	return count > 0
}
