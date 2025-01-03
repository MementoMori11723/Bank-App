package middleware

import (
	"log/slog"
	"net/http"
	"os"
)

var logger *slog.Logger

func init() {
	file, err := os.OpenFile(".server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	logger = slog.New(slog.NewJSONHandler(file, nil))
}

func Log(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		logger.Info("Server info", r.Method, r.URL.Path)
	})
}
