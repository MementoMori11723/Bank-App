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

func Get(args ...any) *sql.Rows {
  db, err := connect()
  defer db.Close()
  if err != nil {
    log.Fatal(err)
    return nil
  }
  rows, err := db.Query(
    "select * from ? where id = ?", 
    args...,
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
    "INSERT INTO ? (name, balance) VALUES (?, ?)",
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
    "UPDATE ? SET balance = ? WHERE name = ?", 
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
    "DELETE FROM ? WHERE name = ?", 
    args...,
  )
  if err != nil {
    log.Fatal(err)
    return
  }
}
