package main

import (
	"fmt"
	"net/http"
	"html/template"
)

var templates *template.Template

func main() {

	fmt.Println("Server GoLang ONLINE na porta: 8080")
	templates = template.Must(template.ParseGlob("templates/*.html"))
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "index.html", nil)

}
