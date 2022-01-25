package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/hello.gohtml"))
	r := mux.NewRouter()
	r.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		vars := make(map[string]string)
		vars["name"] = "World"
		tmpl.Execute(w, vars)
	})
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		tmpl.Execute(w, vars)
	})
	fmt.Print("Start listening on :8000...")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
