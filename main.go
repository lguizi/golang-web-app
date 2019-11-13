package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

var templates *template.Template

func main() {

	fmt.Println("Server GoLang ONLINE na porta: 8080")
	templates = template.Must(template.ParseGlob("templates/*.html"))
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/listar", listarHandler)
	r.HandleFunc("/criar", criarHandler)
	r.HandleFunc("/editar", editarHandler)
	r.HandleFunc("/deletar", deletarHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}


func listarHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "listar.html", nil)
}

func criarHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "criar.html", nil)
}

func editarHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "editar.html", nil)
}

func deletarHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "deletar.html", nil)
}
