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
		"/":       home,
		"/login":  login,
		"/signup": signup,
		"/about":  about,
		"/error":  errorPage,

		"POST /login":               postLogin,
		"POST /signup":              postSignup,
		"POST /details":             postDetails,
		"POST /deposit":             postDeposit,
		"POST /withdraw":            postWithdraw,
		"POST /transfer/{receiver}": postTransfer,
		"POST /history":             postHistory,
		"POST /delete":              postDelete,
	}

	dashboardRoutes = map[string]http.HandlerFunc{
		"/": dashboard,

		"/delete":       deleteFunc,
		"/deposit":      deposit,
		"/withdraw":     withdraw,
		"/transfer":     transfer,
		"/transactions": history,
	}

	pagesDir     = "pages/"
	layout       = pagesDir + "layout.html"
	dashboardDir = pagesDir + "dashboard/"

	baseURL string
)

func Start(port, server_port, server_url string) {
	if server_port == "" && server_url == "" {
		fmt.Println("Port or Server url are is not set!")
		return
	}
  baseURL = "http://localhost:" + server_port + "/"
	if server_url != "" {
    baseURL = server_url
    slog.Info("Server url is set to " + server_url)
	}
	go func() {
		mux := http.NewServeMux()
		dashboard_mux := dashboardMux()
		mux.Handle("/dashboard/", http.StripPrefix("/dashboard", dashboard_mux))
		for route, handler := range routes {
			mux.HandleFunc(route, Handler(handler))
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
