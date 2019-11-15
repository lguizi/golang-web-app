package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var temp = template.Must(template.ParseGlob("templates/*"))

type Cadastros struct {
	Id int
	Nome string
	Tel string
	Email string
}

func bdCad() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "" // Usuário do banco de dados
	dbPass := "" // Senha do banco de dados
	dbName := "" // Nome do banco de dados

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}


func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}

func listar(w http.ResponseWriter, r *http.Request) {
	db := bdCad()
	sel, err := db.Query("select * from cadastros order by id desc;")
	if err != nil {
        	panic(err.Error())
	}

	listaCad := Cadastros{}
	arrayNomes := []Cadastros{}

	for sel.Next() {
		var id int
		var nome, tel, email string
		err = sel.Scan(&id, &nome,  &tel, &email)
		if err != nil {
			panic(err.Error())
		}
		listaCad.Id = id
		listaCad.Nome = nome
		listaCad.Tel = tel
		listaCad.Email = email
        arrayNomes = append(arrayNomes, listaCad)
	}
	temp.ExecuteTemplate(w, "listar.html", arrayNomes)
	defer db.Close()
}

func main() {

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Server ONLINE na porta: 8080")

	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/listar", listar)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
