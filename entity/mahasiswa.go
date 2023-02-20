package entity

type Mahasiswa struct {
	id      int
	nim     int
	nama    string
	email   string
	jurusan string
}

type MahasiswaDTO struct {
	ID      int
	Nim     int
	Nama    string
	Email   string
	Jurusan string
}

func CreateMahasiswa(dto MahasiswaDTO) *Mahasiswa {
	user := &Mahasiswa{
		id:      dto.ID,
		nim:     dto.Nim,
		nama:    dto.Nama,
		email:   dto.Email,
		jurusan: dto.Jurusan,
	}
	return user
}

func (b *Mahasiswa) GetID() int {
	return b.id
}

func (b *Mahasiswa) GetNim() int {
	return b.nim
}

func (b *Mahasiswa) GetNama() string {
	return b.nama
}

func (b *Mahasiswa) GetEmail() string {
	return b.email
}

func (b *Mahasiswa) GetJurusan() string {
	return b.jurusan
}
