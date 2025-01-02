package database

import (
	"bank-app/database/bank"
	"log/slog"
	"net/http"
	"os"
)

var routes = map[string]http.HandlerFunc{
	"/create":       create,
	"/deposit":      deposit,
	"/withdraw":     withdraw,
	"/balance":      balance,
	"/transactions": transactions,
	"/transfer":     transfer,
}

func Server(Port, db_path string) {
	bank.DB_init(db_path)
	mux := http.NewServeMux()
	for route, handler := range routes {
		mux.HandleFunc(route, handler)
	}
	if err := http.ListenAndServe(":"+Port, mux); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
