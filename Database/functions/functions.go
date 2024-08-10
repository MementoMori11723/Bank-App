package functions

import "fmt"

func CreateAccount() string {
  fmt.Println("Create an account")
  return "Account created" //just for testing purposes.
}

func DepositMoney() {
  fmt.Println("Deposit money")
}

func WithdrawMoney() {
  fmt.Println("Withdraw money")
}

func CheckBalance() {
  fmt.Println("Check balance")
}

func ViewTransactionsHistory() {
  fmt.Println("View transactions history")
}

func Settings() {
  fmt.Println("Settings")
}
