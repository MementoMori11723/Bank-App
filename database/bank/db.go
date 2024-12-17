package bank

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log/slog"
	"os"

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

type db_Config struct {
  Database struct {
    Path string `json:"path"` 
  } `json:"database"`
}

var db_path string

func init(){
  file, err := os.ReadFile("./config.json")
  if err != nil {
    slog.Error(err.Error())
  }
  var data db_Config
  err = json.NewDecoder(bytes.NewReader(file)).Decode(&data)
  if err != nil {
    slog.Error(err.Error())
  }
  db_path = data.Database.Path
  if "" == db_path {
    slog.Warn(db_path+" file Doesn't exist! - Creating the file")
    db_path = "./bank.db"
    _, err := os.Create(db_path)
    if err != nil {
      slog.Error(err.Error())
      os.Exit(1)
    }
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
