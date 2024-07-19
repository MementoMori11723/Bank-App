package account

import (
	"database/sql"
	"fmt"
	"math/rand"
)

type Account struct {
	AccountNumber int64
	Name          string
	Balance       float64
	password      int
}

func connectDB() *sql.DB {
	// Connect to database.
	db, err := sql.Open("sqlite3", "./dummy.db")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return db
}

func insertDB(db *sql.DB, user Account) {
	// Insert into database.
}

func fetchDB(accountNumber int32, password int) {
	// Fetch from database.
}

func CreateAccount() {
	var user Account
	var verifyPassword int
	var confirmPassword int
	var err error

	// we need to find a better way to handle the error.
	fmt.Println("Enter your name: ")
  if _,err = fmt.Scanln(&user.Name); err != nil {
    fmt.Println("Error: ", err)
    return
  }

	fmt.Println("Enter your password: ")
	if _, err = fmt.Scanln(&verifyPassword); err != nil {
    fmt.Println("Error: ", err)
    return
  }

	fmt.Println("Confirm your password: ")
	if _, err = fmt.Scanln(&confirmPassword); err != nil {
    fmt.Println("Error: ", err)
    return
  }

	if verifyPassword != confirmPassword {
		fmt.Println("Passwords do not match!")
		return
	} else {
		user.password = verifyPassword
	}

	fmt.Println("Enter your initial deposit: ")
	if _, err = fmt.Scanln(&user.Balance); err != nil {
    fmt.Println("Error: ", err)
    return
  }

	// this model is complicated.
	user.AccountNumber = rand.Int63n(1000000000000000)
	insertDB(connectDB(), user)
	fmt.Println("Account created successfully!")
}

func CheckBalance() {
	var accountNumber int32
	var password int
	fmt.Println("Enter your account number: ")
	fmt.Scanln(&accountNumber)
	fmt.Println("Enter your password: ")
	fmt.Scanln(&password)
	fetchDB(accountNumber, password)
	fmt.Println("Your balance is: ")
}

func DepositMoney() {
	fmt.Println("Money deposited successfully!")
}

func Settings() {
	fmt.Println("Settings:")
}

func ViewTransactionsHistory() {
	fmt.Println("Transactions history:")
}

func WithdrawMoney() {
	fmt.Println("Money withdrawn successfully!")
}
