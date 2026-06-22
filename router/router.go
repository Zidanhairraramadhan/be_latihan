package router

import (
	"be_latihan/config/middleware"
	"be_latihan/handler"
	"be_latihan/model"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Message: "API be_latihan aktif",
		})
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Put("/api/change-password", middleware.JWTProtected(""), handler.ChangePassword)

	mahasiswa := app.Group("/api/mahasiswa")
	mahasiswa.Get("/", middleware.JWTProtected(""), handler.GetAllMahasiswa)
	mahasiswa.Get("/:npm", middleware.JWTProtected("admin"), handler.GetMahasiswaByNPM)
	mahasiswa.Post("/", middleware.JWTProtected("admin"), handler.InsertMahasiswa)
	mahasiswa.Put("/:npm", middleware.JWTProtected("admin"), handler.UpdateMahasiswa)
	mahasiswa.Delete("/:npm", middleware.JWTProtected("admin"), handler.DeleteMahasiswa)
}
