package main

import (
	Cmd "bank-app/cmd"
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
		database.Server()
	}()
	if *port != "" {
		web.Start(*port)
	} else {
		Cmd.Menu()
		fmt.Println("Goodbye")
	}
}
