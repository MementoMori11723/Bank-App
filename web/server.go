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
		"/":      home,
		"/about": about,
		"/error": errorPage,
	}

	dashboardRoutes = map[string]http.HandlerFunc{
		"/": dashboard,
	}
)

func Start(port, server_port string) {
	go func() {
		mux := http.NewServeMux()
		dashboard_mux := dashboardMux()
		mux.Handle("/dashboard/", http.StripPrefix("/dashboard", dashboard_mux))
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

func dashboardMux() *http.ServeMux {
	mux := http.NewServeMux()
	for route, handler := range dashboardRoutes {
		mux.HandleFunc(route, handler)
	}
	return mux
}
