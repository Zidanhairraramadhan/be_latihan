package main

import (
	"be_latihan/config"
	"be_latihan/model"
	"be_latihan/router"
	"log"
	"strings"

	_ "be_latihan/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title API Praktikum 13 - be_latihan
// @version 1.0
// @description Dokumentasi API backend be_latihan menggunakan Golang Fiber, GORM, PostgreSQL, dan JWT.
// @contact.name Praktikum Pemrograman III
// @contact.email praktikum@example.com
// @host 127.0.0.1:3000
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.InitDB()

	// AutoMigrate will construct the tables in the database if they do not exist
	config.GetDB().AutoMigrate(&model.Mahasiswa{}, &model.User{})

	app := fiber.New()

	// CORS configuration to enable frontend access from localhost
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","),
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	router.SetupRoutes(app)

	log.Println("🚀 Server berjalan di http://127.0.0.1:3000")
	log.Fatal(app.Listen(":3000"))
}
