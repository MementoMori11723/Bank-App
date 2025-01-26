package database

import (
	"bank-app/database/bank"
	"bank-app/database/middleware"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
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
	if isPortInUse(Port) {
		slog.Info("Port " + Port + " is already in use. Checking server health...")
		healthURL := fmt.Sprintf("http://localhost:%s/health", Port)
		resp, err := http.Get(healthURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			slog.Info("Server is already running and healthy. Exiting...")
			return
		} else {
			slog.Error("Port is in use, but server is not healthy or /health endpoint is not reachable.", "error", err)
			os.Exit(1)
		}
	}
	go bank.DB_init(db_path)
	mux := http.NewServeMux()

	for route, handler := range routes {
		mux.HandleFunc("POST "+route,
			middleware.Responce(
				handler,
			),
		)
	}

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"status": "ok"}`))
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"status": "error"}`))
		}
	})

	slog.Info("Database - Server started on port " + Port)
	if err := http.ListenAndServe(":"+Port, mux); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func isPortInUse(port string) bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort("localhost", port), 1*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
