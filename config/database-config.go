package config

import (
	"fmt"
	"os"

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/entity" // Import definisi entitas (model) dari aplikasi
	"github.com/joho/godotenv"                              // Import library untuk mengelola variabel lingkungan dari file .env
	"gorm.io/driver/mysql"                                  // Import driver MySQL untuk GORM
	"gorm.io/gorm"                                          // Import library GORM untuk ORM di Go
)

// SetupDatabaseConnection membuat koneksi baru ke database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load() // Load file .env untuk mengambil variabel lingkungan
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// Mendapatkan nilai variabel lingkungan untuk koneksi database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Mengonfigurasi DSN (Data Source Name) untuk koneksi database MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	// Membuat koneksi database menggunakan driver MySQL dan konfigurasi DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	// Menjalankan proses migrasi otomatis untuk tabel-tabel yang didefinisikan dalam aplikasi (entity.Book{} dan entity.User{})
	db.AutoMigrate(&entity.Book{}, &entity.User{})

	return db // Mengembalikan objek koneksi database yang sudah dibuat
}

// CloseDatabaseConnection menutup koneksi database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB() // Mendapatkan objek database SQL dari objek GORM
	if err != nil {
		panic("Failed to close a connection to database")
	}
	dbSQL.Close() // Menutup koneksi database SQL
}
