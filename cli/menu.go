package cli

import (
	"fmt"
	"os"
)

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

	subMenu = map[string]func() []byte{
		"create":       create,
		"deposit":      deposit,
		"withdraw":     withdraw,
		"balance":      balance,
		"transactions": history,
		"transfer":     transfer,
	}
)

func inputFunc[T any](keys []string, m map[string]*T) {
	for _, key := range keys {
		fmt.Print(key, " : ")
		fmt.Scanln(m[key])
	}
}

func errorFunc(err error) {
	fmt.Println("Error: ", err)
	os.Exit(1)
}

func sub_menu(menu string) []byte {
	run, ok := subMenu[menu]
	if !ok {
		panic("Error Occured at sub_menu!")
	}
	data := run()
	return data
}

func Menu(port string) {
  if port == "" {
    fmt.Println("Port is not set!")
    return
  }
  baseURL = "http://localhost:" + port + "/"
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
