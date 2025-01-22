package database

import (
	"bank-app/database/bank"
	"bank-app/database/middleware"
	"log/slog"
	"net/http"
	"os"
)

var routes = map[string]func(*http.Request) (bank.Responce, error){
	"/create":               bank.Create,
	"/deposit":              bank.Deposit,
	"/withdraw":             bank.Withdraw,
	"/transactions":         bank.Transactions,
	"/transfer":             bank.Transfer,
	"/delete":               bank.Delete,
	"/getId":                bank.GetIdByUserName,
	"/details":              bank.Details,
	"/checkUser/{username}": bank.CheckUser,
}

func Server(Port, db_path string) {
	go bank.DB_init(db_path)
	mux := http.NewServeMux()

	for route, handler := range routes {
		mux.HandleFunc("POST "+route,
			middleware.Responce(
				handler,
			),
		)
	}

	slog.Info("Database - Server started on port " + Port)
	if err := http.ListenAndServe(":"+Port, mux); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
