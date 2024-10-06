package db

import (
	"bank-server/config"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func connect() (*sql.DB, error) {
  return sql.Open(
    config.DatabaseDriver, 
    config.DatabaseURL,
  )
}

func Get(tableName string, id int64) *sql.Rows {
  db, err := connect()
  defer db.Close()
  if err != nil {
    log.Fatal(err)
    return nil
  }
  rows, err := db.Query(
    "select * from ? where id = ?", 
    tableName, 
    id,
  )
  if err != nil {
    log.Fatal(err)
    return nil
  }
  return rows
}

func Insert(args ...any) {
  db, err := connect()
  defer db.Close()
  if err != nil {
    log.Fatal(err)
    return
  }
  _, err = db.Exec(
    "INSERT INTO accounts (name, balance) VALUES (?, ?)", 
    args...,
  )
  if err != nil {
    log.Fatal(err)
    return
  }
}

func Update(args ...any) {
  db, err := connect()
  defer db.Close()
  if err != nil {
    log.Fatal(err)
    return
  }
  _, err = db.Exec(
    "UPDATE accounts SET balance = ? WHERE name = ?", 
    args...,
  )
  if err != nil {
    log.Fatal(err)
    return
  }
}

func Delete(args ...any) {
  db, err := connect()
  defer db.Close()
  if err != nil {
    log.Fatal(err)
    return
  }
  _, err = db.Exec(
    "DELETE FROM accounts WHERE name = ?", 
    args...,
  )
  if err != nil {
    log.Fatal(err)
    return
  }
}
