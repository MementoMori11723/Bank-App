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
	serverPort, db_path := config.New()

	port := flag.String("server", "", "port to run the server on")
	flag.Parse()
	fmt.Println("Welcome to the bank")

	go func() {
		database.Server(serverPort, db_path)
	}()

	if *port != "" {
		web.Start(*port)
	} else {
		cli.Menu()
		fmt.Println("Goodbye")
	}
}
