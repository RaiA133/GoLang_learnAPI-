package mapper

import "tugas/entity"

func MapToEntity(listDTO []entity.MahasiswaDTO) []*entity.Mahasiswa {
	var listMahasiswa []*entity.Mahasiswa
	for _, dto := range listDTO {
		mahasiswa := entity.CreateMahasiswa(dto)
		listMahasiswa = append(listMahasiswa, mahasiswa)
	}
	return listMahasiswa
}
