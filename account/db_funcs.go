package account

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func connectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./account/dummy.db")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("Error: ", err)
	}
	return db
}

func insert(user Account) {
	db := connectDB()
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

func fetchAccount(accountNumber int64) (bool,error) {
	var count int
  var err error
	db := connectDB()
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
	db := connectDB()
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
	db := connectDB()
	db.Exec(
		"UPDATE Account SET Balance = ? WHERE AccountNumber = ?",
		user.Balance, user.AccountNumber,
	)
	defer db.Close()
	fmt.Println("Account Balance updated successfully!")
	fmt.Println("Balance: ", user.Balance)
}

func updateName(user Account) {
	db := connectDB()
	db.Exec(
		"UPDATE Account SET Name = ? WHERE AccountNumber = ?",
		user.Name, user.AccountNumber,
	)
	defer db.Close()
	fmt.Println("Account Name updated successfully!")
	fmt.Println("Name: ", user.Name)
}

func updatePassword(user Account) {
	db := connectDB()
	db.Exec(
		"UPDATE Account SET Password = ? WHERE AccountNumber = ?",
		user.password, user.AccountNumber,
	)
	defer db.Close()
	fmt.Println("Account Password updated successfully!")
}
