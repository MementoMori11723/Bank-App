package main

import (
	"fmt"
)

func createAccount() {
	fmt.Println("Account created successfully!")
}

func depositMoney() {
	fmt.Println("Money deposited successfully!")
}

func withdrawMoney() {
	fmt.Println("Money withdrawn successfully!")
}

func checkBalance() {
	fmt.Println("Your balance is: ")
}

func viewTransactionsHistory() {
	fmt.Println("Transactions history:")
}

func settings() {
	fmt.Println("Settings:")
}

func main() {
	fmt.Println("Welcome to dummy bank!")
	fmt.Println("Please select an option:\n1. Create an account\n2. Deposit money\n3. Withdraw money\n4. Check balance\n5. View transactions history\n6. Settings\n7. Exit")
	var signal bool = false
  for !signal{
		var opt int
		fmt.Scanln(&opt)
		switch opt {
		case 1:
			createAccount()
		case 2:
			depositMoney()
		case 3:
			withdrawMoney()
		case 4:
			checkBalance()
		case 5:
			viewTransactionsHistory()
		case 6:
			settings()
		case 7:
			signal = true
		default:
			fmt.Println("Invalid option!")
      signal = true
		}
	}
}
