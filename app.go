package main

import (
	"bank-app/cli"
	"bank-app/config"
	"bank-app/database"
	"bank-app/web"
	"flag"
	"fmt"
)

func main() {
	serverPort, db_path, server_url, close_file := config.New()
	defer close_file()

	web_ui := flag.Bool("web", false, "flag to run the web ui")
	port := flag.String("port", "8001", "port to run the web ui on")
	server_port := flag.String("server-port", serverPort, "port to run the server on")
	flag.Parse()

	fmt.Println("Welcome to Finova")

	if server_url == "" {
		go func() {
			database.Server(*server_port, db_path, server_url)
		}()
	}

	if *web_ui {
		web.Start(*port, *server_port, server_url)
	} else {
		cli.Menu(*server_port, server_url)
		fmt.Println("Goodbye")
	}
}
