package bank

import (
	"bank-app/config"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Account struct {
	ID        int64
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Balance   float64
}

type Transaction struct {
	ID        int
	Sender    string
	Receiver  string
	Amount    float64
	Timestamp string
}

func Connect() (*sql.DB, error) {
	fmt.Println("Connecting to database...")
	db, err := sql.Open("sqlite", config.DB_path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return db, nil
}
