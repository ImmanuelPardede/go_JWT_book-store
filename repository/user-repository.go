package repository

import (
	"log" // Mengimport package log untuk logging

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/entity" // Mengimport package entity untuk model entitas
	"golang.org/x/crypto/bcrypt"                            // Mengimport package bcrypt untuk hashing password
	"gorm.io/gorm"                                          // Mengimport package gorm untuk ORM
)

// UserRepository adalah interface yang mendefinisikan fungsi yang harus diimplementasikan oleh repository User
type UserRepository interface {
	InsertUser(user entity.User) entity.User                    // Fungsi untuk menyimpan user baru
	UpdateUser(user entity.User) entity.User                    // Fungsi untuk mengupdate user
	VerifyCredential(email string, password string) interface{} // Fungsi untuk verifikasi credential user
	IsDuplicateEmail(email string) (tx *gorm.DB)                // Fungsi untuk memeriksa apakah email sudah digunakan
	FindByEmail(email string) entity.User                       // Fungsi untuk mencari user berdasarkan email
	ProfileUser(userID string) entity.User                      // Fungsi untuk mendapatkan profil user berdasarkan ID
}

// userConnection adalah implementasi dari UserRepository
type userConnection struct {
	connection *gorm.DB // Koneksi database menggunakan gorm
}

// NewUserRepository adalah constructor untuk userConnection
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

// InsertUser adalah implementasi fungsi InsertUser dari UserRepository
func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password)) // Menghash password sebelum disimpan
	db.connection.Save(&user)                          // Menyimpan user ke database
	return user                                        // Mengembalikan user yang telah disimpan
}

// UpdateUser adalah implementasi fungsi UpdateUser dari UserRepository
func (db *userConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" { // Jika password diinput, menghash password baru
		user.Password = hashAndSalt([]byte(user.Password))
	} else { // Jika tidak ada input password, menggunakan password yang ada di database
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user) // Menyimpan perubahan data user ke database
	return user               // Mengembalikan user yang telah diupdate
}

// VerifyCredential adalah implementasi fungsi VerifyCredential dari UserRepository
func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user) // Mengambil user berdasarkan email dari database
	if res.Error == nil {                                      // Jika tidak ada error, mengembalikan data user
		return user
	}
	return nil // Jika error, mengembalikan nil
}

// IsDuplicateEmail adalah implementasi fungsi IsDuplicateEmail dari UserRepository
func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user) // Mengambil user berdasarkan email dari database
}

// FindByEmail adalah implementasi fungsi FindByEmail dari UserRepository
func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user) // Mengambil user berdasarkan email dari database
	return user                                         // Mengembalikan data user yang ditemukan
}

// ProfileUser adalah implementasi fungsi ProfileUser dari UserRepository
func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Preload("Books").Preload("Books.User").Find(&user, userID) // Mengambil data user dan relasi Books dari database
	return user                                                              // Mengembalikan profil user yang ditemukan
}

// hashAndSalt adalah fungsi untuk menghash password menggunakan bcrypt
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost) // Menghasilkan hash password dengan cost minimum
	if err != nil {
		log.Print(err)                     // Jika terjadi error, log pesan error
		panic("Failed to hash a password") // Panic jika gagal menghash password
	}
	return string(hash) // Mengembalikan hash password dalam bentuk string
}
