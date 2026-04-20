package db

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func Open(path string) (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, nil, fmt.Errorf("sql.Open: %w", err)
	}

	close := func() {
		db.Close()
	}

	return db, close, nil
}
