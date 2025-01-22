package cli

import (
	"fmt"
	"os"
)

type (
	command    map[string]subCommand
	subCommand struct {
		description string
		run         func() []byte
	}
)

var (
	commands = command{
		"create": subCommand{
			"create account", create,
		},
		"deposit": subCommand{
			"deposit money", deposit,
		},
		"withdraw": subCommand{
			"withdraw money", withdraw,
		},
		"details": subCommand{
			"check details", balance,
		},
		"transfer": subCommand{
			"transfer money", transfer,
		},
		"delete": subCommand{
			"delete account", deleteFunc,
		},
		"transactions": subCommand{
			"check transaction history", history,
		},
	}

	notCommands = map[string]string{
		"help": "show available commands",
		"exit": "exit the program",
	}
)

func inputFunc[T any](keys []string, m map[string]*T) {
	for _, key := range keys {
		fmt.Print(key, ": ")
		fmt.Scanln(m[key])
	}
}

func errorFunc(err error) {
	fmt.Println("Error: ", err)
	os.Exit(1)
}

func sub_menu(menu string) []byte {
	function, ok := commands[menu]
	if !ok {
		panic("Error Occured at sub_menu!")
	}
	data := function.run()
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
			for command, data := range commands {
				fmt.Printf("%-12s - %s\n", command, data.description)
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
