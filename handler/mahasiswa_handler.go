package handler

import (
	"net/http"
	"tugas/repository"
	"tugas/response"
)

func MahasiswaHandler(w http.ResponseWriter, r *http.Request) {

	user, err := repository.GetDataMahasiswa()
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	dataJson, errJson := response.MapResponseListMahasiswa(user)
	if errJson != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(errJson.Error()))
		return
	}

	_, _ = w.Write(dataJson)

}
