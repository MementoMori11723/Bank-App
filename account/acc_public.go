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
	insert(user)
	fmt.Println("Account created successfully!")
}

func CheckBalance() {
  var accountNumber int64
  var password int
  accountNumber = fetchAccountNumber()
  password = fetchPassword()
  Balance := selectBalance(Account{AccountNumber: accountNumber, password: password})
  fmt.Println("Your balance is: ", Balance)
}

func DepositMoney() {
  var user Account
  user.AccountNumber = fetchAccountNumber()
  user.Balance = setBalance()
  updateBalance(user)
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
