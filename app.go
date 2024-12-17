package main

import (
	"bank-app/cli"
	"bank-app/database"
	"bank-app/web"
	"flag"
	"fmt"
)

func main() {
	port := flag.String("server", "", "port to run the server on")
	flag.Parse()
	fmt.Println("Welcome to the bank")
	go func() {
		database.Server("11000")
	}()
	if *port != "" {
		web.Start(*port)
	} else {
		cli.Menu()
		fmt.Println("Goodbye")
	}
}
