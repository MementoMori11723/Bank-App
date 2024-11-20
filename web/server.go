package web

import (
	"embed"
	"fmt"
	"net/http"
)

//go:embed pages/*.html
var pages embed.FS

func Start(port string) {
	go func() {
		http.HandleFunc("/", home)
		http.HandleFunc("/about", about)
		http.HandleFunc("/error", errorPage)
		http.HandleFunc("/dashboard", dashboard)
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
