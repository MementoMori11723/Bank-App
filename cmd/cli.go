package cmd

import (
	"bank-app/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
	}

	fmt.Println("Available commands:")
	for command, description := range commands {
		fmt.Printf("%-12s - %s\n", command, description)
	}
}

func get_response(url string) (string, error) {
	res, err := http.Get("http://localhost:11000/" + url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var bodyRes database.Response
	err = json.Unmarshal(body, &bodyRes)
	if err != nil {
		return "", err
	}
	return bodyRes.Message, nil
}

func fetch_responce(url string) {
	res, err := get_response(url)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}
}

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
