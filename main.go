package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"tugas/handler"

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

	route.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("views/form.html") // Load views/form.html
		temp.Execute(w, nil)                              //

		// if r.Method == http.MethodGet {
		// 	temp, err := template.ParseFiles("views/form.html")
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	temp.Execute(w, nil)
		// } else if r.Method == http.MethodPost {

		// 	r.ParseForm()

		// 	var user entity.MahasiswaDTO
		// 	user.Nama = r.Form.Get("nama")
		// 	user.Nim = r.Form.Get("nim")
		// 	user.Email = r.Form.Get("Email")
		// 	user.Jurusan = r.Form.Get("Jurusan")

		// 	var data = make(map[string]interface{})

		// 	vErrors := validation.Struct(user)

		// 	if vErrors != nil {
		// 		data["pasien"] = pasien
		// 		data["validation"] = vErrors
		// 	} else {
		// 		data["pesan"] = "Data pasien berhasil disimpan"
		// 		pasienModel.Create(pasien)
		// 	}

		// 	temp, _ := template.ParseFiles("views/pasien/add.html")
		// 	temp.Execute(w, data)
		// }
	})

	// http.HandleFunc("/add", repository.PostDataMahasiswa)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", route))

}
