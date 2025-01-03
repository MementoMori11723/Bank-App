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
	if err != nil {
		slog.Error("Database connection failed: " + err.Error())
		return
	}
	defer db.Close()

	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		slog.Error("Failed to set foreign keys: " + err.Error())
		return
	}

	if _, err := db.Exec("PRAGMA journal_mode = WAL;"); err != nil {
		slog.Error("Failed to set journal mode: " + err.Error())
		return
	}
  
	schema, err := os.ReadFile("config/schema.sql")
	if err != nil {
		slog.Error("Failed to read schema file: " + err.Error())
		return
	}

	if _, err := db.Exec(string(schema)); err != nil {
		slog.Error("Failed to execute schema: " + err.Error())
		return
	}
}

func connect() (*sql.DB, error) {
	db, err := sql.Open(
		"sqlite3", db_path,
	)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
