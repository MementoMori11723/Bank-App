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
		"/login": login,
		"/about": about,
		"/error": errorPage,

		"POST /login": postLogin,
	}

	dashboardRoutes = map[string]http.HandlerFunc{
		"/": dashboard,

		"/create":       create,
		"/delete":       deleteFunc,
		"/balance":      balance,
		"/deposit":      deposit,
		"/withdraw":     withdraw,
		"/transfer":     transfer,
		"/transactions": history,
	}

	pagesDir     = "pages/"
	layout       = pagesDir + "layout.html"
	dashboardDir = pagesDir + "dashboard/"

	baseUrl string
)

func Start(port, server_port string) {
	baseUrl = "http://localhost:" + server_port
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
