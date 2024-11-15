package cli

import "fmt"

func Menu() {
	var cmd string
	for true {
		fmt.Print("Enter command: ")
		fmt.Scanln(&cmd)
		switch cmd {
		case "help":
			help()
		case
			"create", "deposit",
			"withdraw", "balance",
			"transactions", "transfer":
			fetch_responce(cmd)
		case "exit":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
