package web

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

//go:embed pages/*.html
var pages embed.FS

func Start(port string) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFS(pages, "pages/layout.html", "pages/home.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			tmpl, err := template.ParseFS(pages, "pages/layout.html", "pages/about.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
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
