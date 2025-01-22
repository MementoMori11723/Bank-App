package web

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var (
	//go:embed pages/*
	pages embed.FS

	routes = map[string]http.HandlerFunc{
		"/":      home,
		"/about": about,
		"/error": errorPage,
	}

	dashboardRoutes = map[string]http.HandlerFunc{
		"/": dashboard,
	}

  pagesDir = "pages/"
  layout = pagesDir+"layout.html"
  dashboardDir = "dashboard/"
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
		fmt.Println("Starting Web Ui server on http://localhost:" + port)
		fmt.Println("Press enter to stop the server...")
		err := client.ListenAndServe()
		if err != nil {
			slog.Error(err.Error())
      fmt.Println("Error starting the server")
      os.Exit(1)
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
