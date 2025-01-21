package web

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
)

var (
	//go:embed pages/*.html
	pages embed.FS

	routes = map[string]http.HandlerFunc{
		"/":          home,
		"/about":     about,
		"/error":     errorPage,
		"/dashboard": dashboard,
	}
)

func Start(port, server_port string) {
	go func() {
		mux := http.NewServeMux()
		for route, handler := range routes {
			mux.HandleFunc(route, handler)
		}
		client := http.Server{
			Addr:    ":" + port,
			Handler: mux,
		}
		fmt.Println("Starting server on http://localhost:" + port)
		fmt.Println("Press enter to stop the server...")
		err := client.ListenAndServe()
		if err != nil {
			slog.Error(err.Error())
			return
		}
	}()
	fmt.Scanln()
}
