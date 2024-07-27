package account

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func getUserName() string {
	fmt.Println("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
		return getUserName()
	}
	name := scanner.Text()
	if name == "" {
		fmt.Println("Name cannot be empty. Please try again.")
		return getUserName()
	}
	return name
}

func handlePassword(password int) int {
  // Having issue with the commmented code below

	// if len(strconv.Itoa(password)) >= 4 {
	// 	fmt.Println("Password must be at least 4 digits. Please try again.")
	// 	return getPassword()
	// }
	fmt.Println("Confirm your password: ")
	var confirmPassword int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	confirmPassword, _ = strconv.Atoi(scanner.Text())
	if password != confirmPassword {
		fmt.Println("Passwords do not match. Please try again.")
		return getPassword()
	}
	return password
}

func fetchPassword() int {
  fmt.Println("Enter your password: ")
  var password int
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  password, _ = strconv.Atoi(scanner.Text())
  return password
}

func fetchAccountNumber() int64 {
  fmt.Println("Enter your account number: ")
  var accountNumber int64
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  accountNumber, _ = strconv.ParseInt(scanner.Text(), 10, 64)
  return accountNumber
}

func getPassword() int {
	fmt.Println("Enter your password: ")
	var password int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	password, _ = strconv.Atoi(scanner.Text())
	return handlePassword(password)
}

func getInitialDeposit() float64 {
	fmt.Println("Enter the amount you want to deposit: ")
	var deposit float64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	deposit, _ = strconv.ParseFloat(scanner.Text(), 64)
	return deposit
}

func generateAccountNumber() int64 {
	var accountNumber string
	for i := 0; i < 16; i++ {
		accountNumber += string(rune(rand.Intn(10)))
	}
	convertedAccountNumber, _ := strconv.ParseInt(accountNumber, 10, 64)
	return convertedAccountNumber
}

func getAccountNumber() int64 {
	AccountNumber := generateAccountNumber()
	condition, err := verifyAccountNumber(AccountNumber)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if condition {
		return getAccountNumber()
	}
	fmt.Println("Here is your account number: ", AccountNumber)
	return AccountNumber
}
