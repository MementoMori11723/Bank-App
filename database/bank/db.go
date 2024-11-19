package bank

import (
	"bank-app/config"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Account struct {
	ID        int64   `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	Balance   float64 `json:"balance"`
}

type Transaction struct {
	ID        int     `json:"id"`
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Amount    float64 `json:"amount"`
	Timestamp string  `json:"timestamp"`
}

func connect() (*sql.DB, error) {
	db, err := sql.Open(
		"sqlite", config.DB_path,
	)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
