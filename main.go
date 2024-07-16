package main

import (
	"fmt"
	"os"

	"bank-cli/account"
)

type Tool func()

func main() {
	fmt.Println("Welcome to dummy bank!")
	fmt.Println("Please select an option:\n1. Create an account\n2. Deposit money\n3. Withdraw money\n4. Check balance\n5. View transactions history\n6. Settings\n7. Exit")

	var signal bool = false

	var tools = []Tool{
		account.CreateAccount,
		account.DepositMoney,
		account.WithdrawMoney,
		account.CheckBalance,
		account.ViewTransactionsHistory,
		account.Settings,
	}

	for !signal {
		var opt int
		_, err := fmt.Scanln(&opt)

		if err != nil {
			fmt.Println("Invalid input!")
			os.Exit(0)
		}

		opt = opt - 1
		if opt != 6 && opt < 6 {
			tools[opt]()
		}

		signal = true
	}
}
