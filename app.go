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
	serverPort, db_path, close_file := config.New()
  defer close_file()

  web_ui := flag.Bool("web", false, "flag to run the web ui")
	port := flag.String("port", "8080", "port to run the web ui on")
  server_port := flag.String("server-port", serverPort, "port to run the server on")
	flag.Parse()

	fmt.Println("Welcome to Finova")

	go func() {
		database.Server(*server_port, db_path)
	}()

	if *web_ui {
		web.Start(*port, *server_port)
	} else {
		cli.Menu(*server_port)
		fmt.Println("Goodbye")
	}
}
