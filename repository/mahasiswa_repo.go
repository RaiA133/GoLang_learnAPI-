package repository

import (
	"tugas/config"
	"tugas/entity"
	"tugas/repository/mapper"
)

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

// func PostDataMahasiswa(w http.ResponseWriter, r *http.Request) {

// if r.Method == http.MethodGet {
// 	temp, err := template.ParseFiles("views/form.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	temp.Execute(w, nil)
// } else if r.Method == http.MethodPost {

// 	r.ParseForm()

// 	var pasien entity.Pasien
// 	pasien.NamaLengkap = r.Form.Get("nama")
// 	pasien.Nim = r.Form.Get("nim")
// 	pasien.Email = r.Form.Get("Email")
// 	pasien.Jurusan = r.Form.Get("Jurusan")

// 	var data = make(map[string]interface{})

// 	vErrors := validation.Struct(pasien)

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
// // mengambil nilai dari form HTML
// nama := r.FormValue("nama")
// nim := r.FormValue("nim")
// email := r.FormValue("email")
// jurusan := r.FormValue("jurusan")

// db, _ := config.MySQLConnection()

// // menyimpan data ke database
// result, err := db.Exec("INSERT INTO mahasiswa (id, nama, nim, email, jurusan) VALUES (?, ?, ?, ?, ?)", "", nama, nim, email, jurusan)
// if err != nil {
// 	http.Error(w, "Gagal menyimpan data ke database", http.StatusInternalServerError)
// 	return
// }
// numRows, _ := result.RowsAffected()
// if numRows == 0 {
// 	http.Error(w, "Tidak ada data yang disimpan", http.StatusInternalServerError)
// 	return
// }
// // mengembalikan respons ke client
// w.WriteHeader(http.StatusCreated)
// w.Write([]byte("Data berhasil disimpan"))

// }
