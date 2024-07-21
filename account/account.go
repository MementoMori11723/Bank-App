package account

import (
	"bufio"
	"fmt"
	"os"
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

func getUserName() string {
	fmt.Println("Enter your name: ")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  return scanner.Text()
}

func handlePassword(password int) int {
  if password < 1000 && password > 9999 {
    fmt.Println("Password must be at least 4 digits. Please try again.")
    return getPassword()
  }
  fmt.Println("Confirm your password: ")
  var confirmPassword int
  fmt.Scanln(&confirmPassword)
  if password != confirmPassword {
    fmt.Println("Passwords do not match. Please try again.")
    return getPassword()
  }
  return password
}

func getPassword() int {
  fmt.Println("Enter your password: ")
  var password int
  fmt.Scanln(&password)
  return handlePassword(password)
}

func getInitialDeposit() float64 {
  fmt.Println("Enter the amount you want to deposit: ")
  var deposit float64
  fmt.Scanln(&deposit)
  return deposit
}

func getAccountNumber() int64 {
  return 1234567890 
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
