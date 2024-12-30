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
	port := flag.String("port", "8000", "port to run the web ui on")
	flag.Parse()

	fmt.Println("Welcome to the bank")

	go func() {
		database.Server(serverPort, db_path)
	}()

	if *web_ui {
		web.Start(*port)
	} else {
		cli.Menu()
		fmt.Println("Goodbye")
	}
}
