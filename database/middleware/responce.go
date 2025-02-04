package middleware

import (
	"bank-app/database/bank"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"Error"`
}

func Responce(next func(*http.Request) (bank.Responce, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Server - Request", r.Method, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Unauthorized"})
			if err != nil {
				slog.Error("Server - Error", "error", err.Error())
			}
			return
		}

		token = token[7:]
		if !validateToken(token) {
			w.WriteHeader(http.StatusUnauthorized)
			err := json.NewEncoder(w).Encode(ErrorResponse{Error: "Unauthorized"})
			if err != nil {
				slog.Error("Server - Error", "error", err.Error())
			}
			return
		}

		data, err := next(r)
		if err != nil {
			slog.Error(err.Error())
		}

		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			err := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			if err != nil {
				slog.Error("Server - Error", "error", err.Error())
			}
		}
	}
}
