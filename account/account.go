package account

import (
	_ "database/sql"
	"fmt"

	"math/rand"
)

type Account struct {
	AccountNumber int64
	Name          string
	Balance       float64
	password      int
}

func insertDB(user Account) {
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

	fmt.Println("Enter your name: ")
	fmt.Scanln(&user.Name)

	fmt.Println("Enter your password: ")
	_, err = fmt.Scanln(&verifyPassword)
	fmt.Println("Confirm your password: ")
	_, err = fmt.Scanln(&confirmPassword)

	if verifyPassword != confirmPassword {
		fmt.Println("Passwords do not match!")
		return
	} else {
		user.password = verifyPassword
	}

	fmt.Println("Enter your initial deposit: ")
	_, err = fmt.Scanln(&user.Balance)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
  
  // Generate a random account number.
  user.AccountNumber = rand.Int63n(1000000000000000)

	insertDB(user)

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
