package account

import (
  "bufio"
  "fmt"
  "os"
)

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
