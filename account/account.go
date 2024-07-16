package account

import (
  "fmt"
  _ "database/sql"
)

type Account struct {
  AccountNumber int
  Name string
  Balance float64
  password string
}

func CheckBalance() {
	fmt.Println("Your balance is: ")
}

func CreateAccount() {
	fmt.Println("Account created successfully!")
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
