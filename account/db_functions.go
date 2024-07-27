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

func insert(user Account) {
	db := connectDB()
	db.Exec(
		"INSERT INTO Account (Name, password, Balance, AccountNumber) VALUES (?, ?)",
		user.Name, user.password, user.Balance, user.AccountNumber,
	)
	defer db.Close()
	fmt.Println("Account inserted successfully!")
	fmt.Println("Name: ", user.Name)
	fmt.Println("Password: ", user.password)
}

func fetch(accountNumber int32, password int) *sql.Rows {
	db := connectDB()
	rows, err := db.Query(
		"SELECT * FROM Account WHERE AccountNumber = ? AND password = ?",
		accountNumber, password,
	)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer rows.Close()
	defer db.Close()
	return rows
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
