package model

type Mahasiswa struct {
	NPM    string `json:"npm" gorm:"column:npm;primaryKey;type:varchar(20)" example:"714240061"`
	Nama   string `json:"nama" gorm:"column:nama;type:varchar(100);not null" example:"Zidan Hairra Ramadhan"`
	Prodi  string `json:"prodi" gorm:"column:prodi;type:varchar(100)" example:"D4 Teknik Informatika"`
	Alamat string `json:"alamat" gorm:"column:alamat;type:varchar(255)" example:"Bandung"`
	Email  string `json:"email" gorm:"column:email;type:varchar(100)" example:"zidan@example.com"`
	NoHP   string `json:"no_hp" gorm:"column:no_hp;type:varchar(20)" example:"081234567890"`
}

func (Mahasiswa) TableName() string { return "mahasiswa" }
