package repository

import (
	"net/http"
	"tugas/config"
	"tugas/entity"
	"tugas/repository/mapper"
)

type mahasiswaModel struct {
}

func GetDataMahasiswa() ([]*entity.Mahasiswa, error) {
	db, err := config.MySQLConnection()
	if err != nil {
		return nil, err
	}

	rows, errRows := db.Query("SELECT * FROM mahasiswa")
	if errRows != nil {
		return nil, errRows
	}

	var result []entity.MahasiswaDTO
	for rows.Next() {
		var each = entity.MahasiswaDTO{}
		errScan := rows.Scan(&each.ID, &each.Nim, &each.Nama, &each.Email, &each.Jurusan)
		if errScan != nil {
			return nil, errScan
		}
		result = append(result, each)
	}
	return mapper.MapToEntity(result), nil
}

func PostDataMahasiswa(w http.ResponseWriter, r *http.Request) {
	db, _ := config.MySQLConnection()
	if r.Method == "POST" {
		nama := r.FormValue("nama")
		nim := r.FormValue("nim")
		email := r.FormValue("email")
		jurusan := r.FormValue("jurusan")
		insForm, err := db.Prepare("INSERT INTO mahasiswa(nama, nim, email, jurusan) VALUES(?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(nama, nim, email, jurusan)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}
