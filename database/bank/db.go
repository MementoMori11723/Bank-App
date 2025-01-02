package bank

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db_path string

func DB_init(path string) {
	if path == "" {
		path = "bank.db"
	}

	db_path = path

	db, err := connect()
	defer db.Close()
	if err != nil {
		slog.Error(err.Error())
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		slog.Error(err.Error())
	}

	if _, err := db.Exec("PRAGMA journal_mode = WAL;"); err != nil {
		slog.Error(err.Error())
	}

	schema, err := os.ReadFile("../../config/schema.sql")
	if err != nil {
		slog.Error(err.Error())
	}

	schemaString := string(schema)
	if _, err := db.Exec(schemaString); err != nil {
		slog.Error(err.Error())
	}
}

func connect() (*sql.DB, error) {
	db, err := sql.Open(
		"sqlite", db_path,
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
