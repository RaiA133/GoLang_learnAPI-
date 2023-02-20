package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"tugas/handler"
	"tugas/repository"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// http://localhost:8080/
	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("views/index.html") // Load views/index.html
		temp.Execute(w, nil)                               //
	}).Methods(http.MethodGet)

	// http://localhost:8080/user
	route.HandleFunc("/user", handler.MahasiswaHandler).Methods(http.MethodGet)

	route.HandleFunc("/form", handler.AddData)
	route.HandleFunc("/add", repository.PostDataMahasiswa)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", route))

}
