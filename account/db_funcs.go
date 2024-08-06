// Description: This file contains all the functions that interact with the database.

// there are 4 sections in this file:
// 1. Connection functions
// 2. Insert functions
// 3. Update functions
// 4. Select functions

package account

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Connection functions section

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./account/dummy.db")
	return db, err
}

// Insert functions section

func insert(user Account) {
	db, _ := connectDB()
	_, err := db.Exec(
		"INSERT INTO Account (Name, password, Balance, AccountNumber) VALUES (?, ?, ?, ?)",
		user.Name, user.password, user.Balance, user.AccountNumber,
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer db.Close()
	fmt.Println("Account inserted successfully!")
	fmt.Println("Name: ", user.Name)
	fmt.Println("Password: ", user.password)
}

// Update functions section

func updateDB(user Account, query string, condition string) string {
	db, err := connectDB()
	defer db.Close()
	if err != nil {
		return err.Error()
	}
	switch condition {
	case "Name":
		db.Exec(query, user.Name, user.AccountNumber)
	case "Password":
		db.Exec(query, user.password, user.AccountNumber)
	case "Balance":
		db.Exec(query, user.Balance, user.AccountNumber)
	default:
		return "Invalid Condition"
	}
	return "Updated Data without issues!"
}

func updateBalance(user Account) {
	msg := updateDB(user, "UPDATE Account SET Balance = ? WHERE AccountNumber = ?", "Balance")
	fmt.Println(msg)
}

func updateName(user Account) {
	msg := updateDB(user, "UPDATE Account SET Name = ? WHERE AccountNumber = ?", "Name")
	fmt.Println(msg)
}

func updatePassword(user Account) {
	msg := updateDB(user, "UPDATE Account SET password = ? WHERE AccountNumber = ?", "Password")
	fmt.Println(msg)
}

// Select functions section

func fetch(user Account, query string, condition string) (*sql.Rows, error) {
	db, err := connectDB()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	switch condition {
	case "AccountNumber":
    row, err := db.Query(query, user.AccountNumber)
		if err != nil {
			return row, err
		}
    return row, nil
	case "Balance":
    row, err := db.Query(query, user.AccountNumber, user.password)
		if err != nil {
			return row, err
		}
    return row, nil
	default:
		return nil, errors.New("Invalid Condition")
	}
}

func selectAndVerifyAccount(accountNumber int64) (bool, error) {
  user := Account{AccountNumber: accountNumber}
  rows,err := fetch(user, "SELECT AccountNumber FROM Account WHERE AccountNumber = ?", "AccountNumber")
  defer rows.Close()
  if err != nil {
    return false, err
  }
  for rows.Next() {
    rows.Scan(&user.AccountNumber)
    if user.AccountNumber == accountNumber {
      return true, nil
    }
  }
  return false, nil
}

func selectBalance(user Account) float64 {
  rows,err := fetch(user, "SELECT Balance FROM Account WHERE AccountNumber = ? AND password = ?", "Balance")
  defer rows.Close()
  if err != nil {
    fmt.Println("Error: ", err)
  }
  for rows.Next() {
    rows.Scan(&user.Balance)
  }
  return user.Balance
}
