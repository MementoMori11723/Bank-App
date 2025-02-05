package main

import (
	"bank-app/cli"
	"bank-app/config"
	"bank-app/database"
	"bank-app/web"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	serverPort, db_path, server_url, close_file := config.New()
	defer close_file()

	web_ui := flag.Bool("web", false, "flag to run the web ui")
	port := flag.String("port", "8001", "port to run the web ui on")
	server_port := flag.String("server-port", serverPort, "port to run the server on")
	flag.Parse()

	if server_url != "" {
		slog.Info("Server is already running on " + server_url + " Check if it is running")
		if res, err := http.Get(server_url + "/health"); err != nil || res.StatusCode != http.StatusOK {
			fmt.Println("Server is not running or not reachable")
			slog.Error("Server is not running or not reachable", "Error", err)
      slog.Info("Starting the local server")
      server_url = ""
		}
	}

	go func() {
		database.Server(*server_port, db_path, server_url)
	}()

	fmt.Println("Welcome to Finova")
	if *web_ui {
		web.Start(*port, *server_port, server_url)
	} else {
		cli.Menu(*server_port, server_url)
		fmt.Println("Goodbye")
	}
}
