package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file tidak ditemukan, menggunakan environment variable")
	}

	dsn := os.Getenv("SUPABASE_DSN")
	if dsn == "" {
		log.Fatal("❌ SUPABASE_DSN tidak ditemukan di .env")
	}

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal koneksi ke database:", err)
	}

	db = database
	log.Println("✅ Database terkoneksi")
}

func GetDB() *gorm.DB {
	return db
}

func GetAllowedOrigins() []string {
	origins := os.Getenv("ALLOWED_ORIGINS")
	if origins == "" {
		return []string{"http://localhost:5173", "http://localhost:3000", "http://amused-fulfillment-production-6e6d.up.railway.app"}
	}
	return strings.Split(origins, ",")
}
