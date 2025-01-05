package middleware

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"Error"`
}

func Responce[T any](next func(*http.Request) (T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data, err := next(r)
		if err != nil {
			slog.Error(err.Error())
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(
				ErrorResponse{
					Error: err.Error(),
				},
			)
			if err != nil {
				slog.Error(err.Error())
			}
		}
	}
}
