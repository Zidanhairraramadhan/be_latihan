package model

// Response adalah struct response umum yang dipakai hampir di semua endpoint
type Response struct {
	Message string      `json:"message" example:"detail pesan"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty" example:"detail error"`
}

// Response200 adalah struct contoh response sukses 200 OK
type Response200 struct {
	Message string      `json:"message" example:"berhasil mengambil data"`
	Data    interface{} `json:"data,omitempty"`
}

// Response201 adalah struct contoh response sukses 201 Created
type Response201 struct {
	Message string      `json:"message" example:"data berhasil ditambahkan"`
	Data    interface{} `json:"data,omitempty"`
}

// Response401 adalah struct contoh response 401 Unauthorized
// Dikirim ketika token JWT tidak ada atau tidak valid
type Response401 struct {
	Message string `json:"message" example:"token tidak valid atau tidak ditemukan"`
}

// Response403 adalah struct contoh response 403 Forbidden
// Dikirim ketika user sudah login tetapi role-nya tidak memiliki akses ke endpoint ini
type Response403 struct {
	Message string `json:"message" example:"user tidak memiliki akses untuk fitur ini"`
}
