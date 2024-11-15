package web

import (
	"fmt"
	"net/http"
)

func Start(port string) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "<h1>Hello, you've requested: %s\n</h1>", r.URL.Path)
		})
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "<h1>This is an example of a simple HTTP server</h1>")
		})
		fmt.Println("Starting server on http://localhost:" + port)
		fmt.Println("Press enter to stop the server...")
		err := http.ListenAndServe(":"+port, nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
			return
		}
	}()
	fmt.Scanln()
}
