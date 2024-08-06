// Functions present in this file

// connectionDB - returns a pointer to a connection variable to the database
// insert - takes in Account varible and inserts it to the database 
// fetchAccount - takes accountNumber and returns a boolean and an error variable
// fetchBalance - takes accountNumber and returns a float
// updateBalance - takes in Account varible and updates Balance
// updateName - takes in Account varible and updates Name
// updatePassword - takes in Account varible and updates Password

package account

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func connectDB() (*sql.DB,error) {
	db, err := sql.Open("sqlite3", "./account/dummy.db")
	return db,err
}

func insert(user Account) {
	db,_ := connectDB()
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

func updateDB(user Account, query string) (string) {
  db,err := connectDB()
  defer db.Close()
  if err != nil {
    return err.Error()
  }
  return "Updated Data without issues!"
}

func fetchAccount(accountNumber int64) (bool,error) {
	var count int
  var err error
	db,_ := connectDB()
	defer db.Close()
	rows, err := db.Query("SELECT COUNT(AccountNumber) FROM Account WHERE AccountNumber = ?", accountNumber)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
	}
	if count == 0 {
		return true, err
	}
	return false, err
}

func fetchBalance(accountNumber int64, password int) float64 {
	db,_ := connectDB()
	var Balance float64
	rows, err := db.Query(
		"SELECT Balance FROM Account WHERE AccountNumber = ? AND password = ?",
		accountNumber, password,
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&Balance)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
	defer db.Close()
	return Balance
}

func updateBalance(user Account) {
	db,_ :=connectDB()
	db.Exec(
		"UPDATE Account SET Balance = ? WHERE AccountNumber = ?",
		user.Balance, user.AccountNumber,
	)
	defer db.Close()
	fmt.Println("Account Balance updated successfully!")
	fmt.Println("Balance: ", user.Balance)
}

func updateName(user Account) {
	db,_ :=connectDB()
	db.Exec(
		"UPDATE Account SET Name = ? WHERE AccountNumber = ?",
		user.Name, user.AccountNumber,
	)
	defer db.Close()
	fmt.Println("Account Name updated successfully!")
	fmt.Println("Name: ", user.Name)
}

func updatePassword(user Account) {
	db,_ :=connectDB()
	db.Exec(
		"UPDATE Account SET Password = ? WHERE AccountNumber = ?",
		user.password, user.AccountNumber,
	)
	defer db.Close()
	fmt.Println("Account Password updated successfully!")
}
