// This is Level - 1 also the main file.

// It will provide an interface to the user by using http server to host and process the requests.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/MementoMori11723/Bank-App/database/functions"
)

type Data struct {
	Message string `json:"message"`
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	msg := functions.CreateAccount()
	data := Data{Message: msg}
	json.NewEncoder(w).Encode(data)
}

func main() {
	fmt.Println("Welcome to Dummy Bank!")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	http.HandleFunc("/", createAccount)
	http.ListenAndServe(":"+PORT, nil)
}
