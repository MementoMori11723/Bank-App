package web

import (
	"context"
	"embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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
	var wg sync.WaitGroup

	if server_port == "" && server_url == "" {
		fmt.Println("Port or Server url are is not set!")
		return
	}

	baseURL = "http://localhost:" + server_port + "/"
	if server_url != "" {
		baseURL = server_url
		slog.Info("Server url is set to " + server_url)
	}

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

	wg.Add(1)
	go startServer(&wg, port, &client)

  stop := make(chan os.Signal, 1)
  signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
  enterPressed := make(chan bool)

  go func() {
    fmt.Println("Press enter to stop the server...")
    fmt.Scanln()
    enterPressed <- true
  }()

  select {
  case <-stop:
  case <-enterPressed:
  }
  fmt.Println("\nShutting down the server...")
  
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()
  if err := client.Shutdown(ctx); err != nil {
    slog.Error(err.Error())
  } else {
    fmt.Println("Server gracefully stopped")
  }
  wg.Wait()
}

func startServer(wg *sync.WaitGroup, port string, client *http.Server) {
  defer wg.Done()
	fmt.Println("Starting Web Ui server on http://localhost:" + port)
	
  err := client.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error(err.Error())
		fmt.Println("Error starting the server")
		os.Exit(1)
	}
}

func dashboardMux() *http.ServeMux {
	mux := http.NewServeMux()
	for route, handler := range dashboardRoutes {
		mux.HandleFunc(route, handler)
	}
	return mux
}
