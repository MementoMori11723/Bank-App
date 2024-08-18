// This is Level - 1 also the main file.

// it will provide an interface to the user by using http server to host and process the requests.

package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	// "github.com/mementomori11723/bank-app/database/functions/process"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
  })
	http.ListenAndServe(":"+PORT, nil)
}
