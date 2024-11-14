package cmd

import (
	"bank-app/database"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func help() {
  fmt.Println("Available commands:")
  fmt.Println("help - show this help")
  fmt.Println("create - create a new account")
  fmt.Println("deposit - deposit money")
  fmt.Println("withdraw - withdraw money")
  fmt.Println("balance - check balance")
  fmt.Println("transactions - check transactions")
  fmt.Println("transfer - transfer money")
  fmt.Println("exit - exit the application")
}

func get_response(url string) (string, error) {
  res, err := http.Get("http://localhost:11000/"+url)
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
    case "create":
      fetch_responce(cmd)
    case "deposit":
      fetch_responce(cmd)
    case "withdraw":
      fetch_responce(cmd)
    case "balance":
      fetch_responce(cmd)
    case "transactions":
      fetch_responce(cmd)
    case "transfer":
      fetch_responce(cmd)
    case "exit":
      return
    default:
      fmt.Println("Invalid command")
    }
  }
}
