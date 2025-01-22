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
		w.WriteHeader(http.StatusOK)
		slog.Info("Server - Request", r.Method, r.URL.Path)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
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
