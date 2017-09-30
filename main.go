package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080", nil)
}

var t *template.Template

func init() {
	fmt.Println("Initialize")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	t, _ = template.ParseFiles("views/home.html", "views/layout.html")
	err := t.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
