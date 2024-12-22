package cli

import "fmt"

type command map[string]string

var (
	commands = command{
		"create":       "create a new account",
		"deposit":      "deposit money",
		"withdraw":     "withdraw money",
		"balance":      "check balance",
		"transactions": "check transactions",
		"transfer":     "transfer money",
	}

	notCommands = command{
		"help": "show this help",
		"exit": "exit the application",
	}
)

func Menu() {
	var cmd string
	for true {
		fmt.Print("Enter command: ")
		fmt.Scanln(&cmd)
		_, ok := commands[cmd]
		if cmd != "exit" && !ok {
			fmt.Println("Invalid command!\nAvailable commands:")
			for command, description := range notCommands {
				fmt.Printf("%-12s - %s\n", command, description)
			}
			for command, description := range commands {
				fmt.Printf("%-12s - %s\n", command, description)
			}
		} else {
			if ok {
				fetch_responce(cmd)
			} else {
				return
			}
		}
	}
}
