package database

import (
	"net/http"
)

func Server() {
  http.HandleFunc("/create", create)
  http.HandleFunc("/deposit", deposit)
  http.HandleFunc("/withdraw", withdraw)
  http.HandleFunc("/balance", balance)
  http.HandleFunc("/transactions", transactions)
  http.HandleFunc("/transfer", transfer)
  http.ListenAndServe(":11000", nil)
}
