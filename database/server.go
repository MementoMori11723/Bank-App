package database

import (
	"bank-app/database/bank"
	"bank-app/database/middleware"
	"encoding/json"
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

func Server(Port, db_path, server_url string) {
	if isPortInUse(Port) {
		slog.Info("Port " + Port + " is already in use. Checking server health...")
		healthURL := fmt.Sprintf("http://localhost:%s/health", Port)
		resp, err := http.Get(healthURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			slog.Info("Server is already running and healthy. Exiting...")
			middleware.BaseURL("http://localhost:" + Port)
			return
		} else {
			slog.Error("Port is in use, but server is not healthy or /health endpoint is not reachable.", "error", err)
			os.Exit(1)
		}
	}

	if server_url != "" {
		slog.Info("Server URL is set to " + server_url + ". Checking server health...")
		healthURL := fmt.Sprintf("%s/health", server_url)
		resp, err := http.Get(healthURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			slog.Info("Server is already running and healthy. Exiting...")
			middleware.BaseURL(server_url)
			return
		} else {
			slog.Error("Server URL is set, but server is not healthy or /health endpoint is not reachable.", "error", err)
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
