package repository

import (
	"be_latihan/config"
	"be_latihan/model"
)

func GetAllMahasiswa() ([]model.Mahasiswa, error) {
	var mahasiswas []model.Mahasiswa
	result := config.GetDB().Find(&mahasiswas)
	return mahasiswas, result.Error
}

func GetMahasiswaByNPM(npm string) (model.Mahasiswa, error) {
	var mahasiswa model.Mahasiswa
	result := config.GetDB().First(&mahasiswa, "npm = ?", npm)
	return mahasiswa, result.Error
}

func InsertMahasiswa(mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error) {
	result := config.GetDB().Create(mahasiswa)
	return mahasiswa, result.Error
}

func UpdateMahasiswa(npm string, mahasiswa *model.Mahasiswa) (*model.Mahasiswa, error) {
	var existing model.Mahasiswa
	result := config.GetDB().First(&existing, "npm = ?", npm)
	if result.Error != nil {
		return nil, result.Error
	}

	existing.Nama = mahasiswa.Nama
	existing.Prodi = mahasiswa.Prodi
	existing.Alamat = mahasiswa.Alamat
	existing.Email = mahasiswa.Email
	existing.NoHP = mahasiswa.NoHP

	result = config.GetDB().Save(&existing)
	return &existing, result.Error
}

func DeleteMahasiswa(npm string) error {
	result := config.GetDB().Where("npm = ?", npm).Delete(&model.Mahasiswa{})
	return result.Error
}
