package cli

import "fmt"

func help() {
	commands := map[string]string{
		"help":         "show this help",
		"create":       "create a new account",
		"deposit":      "deposit money",
		"withdraw":     "withdraw money",
		"balance":      "check balance",
		"transactions": "check transactions",
		"transfer":     "transfer money",
		"exit":         "exit the application",
		"-server":      "run the application as a web server on the specified port",
	}

	fmt.Println("Available commands:")
	for command, description := range commands {
		fmt.Printf("%-12s - %s\n", command, description)
	}
}
