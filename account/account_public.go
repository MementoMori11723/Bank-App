package account

import (
	"fmt"
)

type Account struct {
	AccountNumber int64
	Name          string
	Balance       float64
	password      int
}

func CreateAccount() {
  var user Account
  user.Name = getUserName()
  user.password = getPassword()
  user.Balance = getInitialDeposit()
  user.AccountNumber = getAccountNumber()

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
