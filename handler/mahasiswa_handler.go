package handler

import (
	"be_latihan/model"
	"be_latihan/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllMahasiswa godoc
// @Summary Ambil semua data mahasiswa
// @Description Mengambil seluruh data mahasiswa. Endpoint ini membutuhkan token JWT, tetapi tidak membatasi role admin.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} model.Response200 "Berhasil mengambil semua data mahasiswa"
// @Failure 401 {object} model.Response401 "Token JWT tidak ada atau tidak valid"
// @Failure 500 {object} model.Response "Terjadi kesalahan pada server"
// @Router /api/mahasiswa/ [get]
func GetAllMahasiswa(c *fiber.Ctx) error {
	mahasiswas, err := repository.GetAllMahasiswa()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengambil data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "berhasil mengambil data mahasiswa",
		Data:    mahasiswas,
	})
}

// GetMahasiswaByNPM godoc
// @Summary Ambil data mahasiswa berdasarkan NPM
// @Description Mengambil satu data mahasiswa berdasarkan NPM. Endpoint ini hanya dapat diakses oleh role admin.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path string true "NPM mahasiswa"
// @Success 200 {object} model.Response200 "Data mahasiswa berhasil ditemukan"
// @Failure 400 {object} model.Response "Format NPM tidak valid"
// @Failure 401 {object} model.Response401 "Token JWT tidak ada atau tidak valid"
// @Failure 403 {object} model.Response403 "Role user tidak memiliki akses ke endpoint ini"
// @Failure 404 {object} model.Response "Data mahasiswa tidak ditemukan"
// @Failure 500 {object} model.Response "Terjadi kesalahan pada server"
// @Router /api/mahasiswa/{npm} [get]
func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm := c.Params("npm")
	mahasiswa, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(model.Response{
				Message: "mahasiswa tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengambil data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "data mahasiswa ditemukan",
		Data:    mahasiswa,
	})
}

// InsertMahasiswa godoc
// @Summary Tambah data mahasiswa
// @Description Menambahkan data mahasiswa baru. Endpoint ini hanya dapat diakses oleh role admin.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body model.Mahasiswa true "Payload data mahasiswa"
// @Success 201 {object} model.Response201 "Mahasiswa berhasil ditambahkan"
// @Failure 400 {object} model.Response "Payload tidak valid atau NPM/Nama kosong"
// @Failure 401 {object} model.Response401 "Token JWT tidak ada atau tidak valid"
// @Failure 403 {object} model.Response403 "Role user tidak memiliki akses ke endpoint ini"
// @Failure 500 {object} model.Response "Terjadi kesalahan pada server"
// @Router /api/mahasiswa/ [post]
func InsertMahasiswa(c *fiber.Ctx) error {
	var mahasiswa model.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	if mahasiswa.NPM == "" || mahasiswa.Nama == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "npm dan nama wajib diisi",
		})
	}

	data, err := repository.InsertMahasiswa(&mahasiswa)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal menambahkan mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Message: "mahasiswa berhasil ditambahkan",
		Data:    data,
	})
}

// UpdateMahasiswa godoc
// @Summary Ubah data mahasiswa
// @Description Mengubah data mahasiswa berdasarkan NPM. NPM pada body akan dipaksa mengikuti NPM pada URL.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path string true "NPM mahasiswa"
// @Param request body model.Mahasiswa true "Payload data mahasiswa"
// @Success 200 {object} model.Response200 "Data mahasiswa berhasil diubah"
// @Failure 400 {object} model.Response "Format NPM atau payload tidak valid"
// @Failure 401 {object} model.Response401 "Token JWT tidak ada atau tidak valid"
// @Failure 403 {object} model.Response403 "Role user tidak memiliki akses ke endpoint ini"
// @Failure 404 {object} model.Response "Data mahasiswa tidak ditemukan"
// @Failure 500 {object} model.Response "Terjadi kesalahan pada server"
// @Router /api/mahasiswa/{npm} [put]
func UpdateMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")
	var mahasiswa model.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	data, err := repository.UpdateMahasiswa(npm, &mahasiswa)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(model.Response{
				Message: "mahasiswa tidak ditemukan",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengupdate mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "mahasiswa berhasil diupdate",
		Data:    data,
	})
}

// DeleteMahasiswa godoc
// @Summary Hapus data mahasiswa
// @Description Menghapus data mahasiswa berdasarkan NPM. Endpoint ini hanya dapat diakses oleh role admin.
// @Tags Mahasiswa
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param npm path string true "NPM mahasiswa"
// @Success 200 {object} model.Response200 "Mahasiswa berhasil dihapus"
// @Failure 400 {object} model.Response "Format NPM tidak valid"
// @Failure 401 {object} model.Response401 "Token JWT tidak ada atau tidak valid"
// @Failure 403 {object} model.Response403 "Role user tidak memiliki akses ke endpoint ini"
// @Failure 500 {object} model.Response "Terjadi kesalahan pada server"
// @Router /api/mahasiswa/{npm} [delete]
func DeleteMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")
	err := repository.DeleteMahasiswa(npm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal menghapus mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "mahasiswa berhasil dihapus",
	})
}
