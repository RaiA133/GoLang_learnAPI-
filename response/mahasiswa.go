package response

import (
	"encoding/json"
	"tugas/entity"
	_interface "tugas/entity/interface"
)

type JSONListMahasiswa struct {
	ID      int    `json:"id"`
	Nim     int    `json:"nim"`
	Nama    string `json:"nama"`
	Email   string `json:"email"`
	Jurusan string `json:"jurusan"`
}

func MapResponseListMahasiswa(listMahasiswa []*entity.Mahasiswa) ([]byte, error) {
	var dataJson []*JSONListMahasiswa
	for _, mahasiswa := range listMahasiswa {
		var jsonSingle = EntityToJson(mahasiswa)

		dataJson = append(dataJson, jsonSingle)
	}
	byteJson, err := json.Marshal(dataJson)
	if err != nil {
		return nil, err
	}

	return byteJson, nil
}

func EntityToJson(user _interface.MahasiswaInterfaceGet) *JSONListMahasiswa {
	jsonSingle := &JSONListMahasiswa{
		ID:      user.GetID(),
		Nim:     user.GetNim(),
		Nama:    user.GetNama(),
		Email:   user.GetEmail(),
		Jurusan: user.GetJurusan(),
	}

	return jsonSingle
}
