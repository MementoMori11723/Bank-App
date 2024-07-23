package account

import (
  "database/sql"
  "fmt"

  _ "github.com/mattn/go-sqlite3"
)

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./account/dummy.db")
  defer db.Close()
	
  if err != nil {
		fmt.Println("Error: ", err)
	}

  if err = db.Ping(); err != nil {
    fmt.Println("Error: ", err)
  }

	return db
}

func verifyAccountNumber(accountNumber int64) (bool, error) {
  // Verify account number.
  if accountNumber != 0 {
    return true, nil
  }
  return false, nil
}

func insertDB(user Account) {
  db := connectDB()
  db.Exec(
    "INSERT INTO account (Name, password, Balance, AccountNumber) VALUES (?, ?)", 
    user.Name, user.password, user.Balance, user.AccountNumber,
  )
  defer db.Close()
	fmt.Println("Account inserted successfully!")
  fmt.Println("Name: ", user.Name)
  fmt.Println("Password: ", user.password)
}

func fetchDB(accountNumber int32, password int) {
	// Fetch from database.
}
