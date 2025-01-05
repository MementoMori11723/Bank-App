package database

import (
	"bank-app/database/bank"
	"bank-app/database/middleware"
	"log/slog"
	"net/http"
	"os"
)

var routes = map[string]func(*http.Request) (bank.Responce, error){
	"/create":       bank.Create,
	"/deposit":      bank.Deposit,
	"/withdraw":     bank.Withdraw,
	"/balance":      bank.Balance,
	"/transactions": bank.Transactions,
	"/transfer":     bank.Transfer,
	"/delete":       bank.Delete,
}

func Server(Port, db_path string) {
	go bank.DB_init(db_path)
	mux := http.NewServeMux()

	for route, handler := range routes {
		mux.HandleFunc(route,
			middleware.Responce(
				handler,
			),
		)
	}

	if err := http.ListenAndServe(":"+Port, mux); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
