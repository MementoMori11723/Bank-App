package database

import (
	"bank-app/database/bank"
	"bank-app/database/middleware"
	"encoding/json"
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

func Server(Port, db_path, server_url string) {
	if isPortInUse(Port) {
		healthCheckURL := "http://localhost:" + Port + "/health"
    slog.Info("Checking if server is already running")
		res, err := http.Get(healthCheckURL)
		if err == nil && res.StatusCode == http.StatusOK {
			return
		}
    slog.Error(err.Error())
	}

	bank.DB_init(db_path)
	mux := http.NewServeMux()
  slog.Info("Server is starting")

	for route, handler := range routes {
		mux.HandleFunc("POST "+route,
			middleware.Responce(
				handler,
			),
		)
	}

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("X-Request-Type") == "secret" {
			err := json.NewEncoder(w).
				Encode(middleware.SecretKeyResponse{
					Key: middleware.GetSecretKey(),
				})
			if err != nil {
				slog.Error(err.Error())
				http.Error(w, "Failed to Get Secret Key!", http.StatusInternalServerError)
			}
			return
		}
		err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
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
