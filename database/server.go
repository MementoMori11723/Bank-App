package database

import (
	"log/slog"
	"net/http"
	"os"
)

var (
  routes = map[string]http.HandlerFunc{
    "/create": create,
    "/deposit": deposit,
    "/withdraw":withdraw,
    "/balance": withdraw, 
    "/transactions": transactions,
    "/transfer": transfer,
  }
)

func Server(Port string) {
  mux := http.NewServeMux()
  for route, handler := range routes {
    mux.HandleFunc(route, handler)
  }
  if err := http.ListenAndServe(":" + Port, mux); 
  err != nil {
    slog.Error(err.Error())
    os.Exit(1) 
  }
}
