package web

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	tmpl, err := template.ParseFS(pages, layout, pagesDir+"home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, pagesDir+"about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, pagesDir+"login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func errorPage(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	tmpl, err := template.ParseFS(pages, layout, pagesDir+"error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/error", http.StatusFound)
	}
	tmpl, err := template.ParseFS(pages, layout, dashboardDir+"index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func signup(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, pagesDir+"create.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func deleteFunc(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, dashboardDir+"delete.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func deposit(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, dashboardDir+"deposit.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func withdraw(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, dashboardDir+"withdraw.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func history(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, dashboardDir+"history.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func transfer(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFS(pages, layout, dashboardDir+"transfer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
