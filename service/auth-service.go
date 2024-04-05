package service

import (
	"log" // Mengimport package log untuk logging

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/dto"        // Mengimport package dto untuk DTO (Data Transfer Object)
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/entity"     // Mengimport package entity untuk model entitas
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/repository" // Mengimport package repository untuk interaksi dengan database
	"github.com/mashingan/smapping"                             // Mengimport package smapping untuk mapping struct
	"golang.org/x/crypto/bcrypt"                                // Mengimport package bcrypt untuk hashing password
)

// AuthService adalah interface yang mendefinisikan fungsi yang harus diimplementasikan oleh service Auth
type AuthService interface {
	VerifyCredential(email string, password string) interface{} // Fungsi untuk verifikasi credential user
	CreateUser(user dto.RegisterDTO) entity.User                // Fungsi untuk membuat user baru
	FindByEmail(email string) entity.User                       // Fungsi untuk mencari user berdasarkan email
	IsDuplicateEmail(email string) bool                         // Fungsi untuk memeriksa apakah email sudah digunakan
}

// authService adalah implementasi dari AuthService
type authService struct {
	userRepository repository.UserRepository // Menggunakan repository untuk interaksi dengan database user
}

// NewAuthService adalah constructor untuk authService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

// VerifyCredential adalah implementasi fungsi VerifyCredential dari AuthService
func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password) // Memanggil repository untuk verifikasi credential
	if v, ok := res.(entity.User); ok {                             // Memeriksa apakah hasil verifikasi adalah instance dari entity.User
		comparePassword := comparePassword(v.Password, []byte(password)) // Membandingkan password yang dihash dengan password input
		if v.Email == email && comparePassword {                         // Jika email dan password cocok
			return res // Mengembalikan hasil verifikasi
		}
		return false // Jika tidak cocok, mengembalikan false
	}
	return res // Mengembalikan hasil verifikasi
}

// CreateUser adalah implementasi fungsi CreateUser dari AuthService
func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}                                        // Mendeklarasikan variabel untuk menyimpan data user yang akan dibuat
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user)) // Mengisi struct user dengan data dari DTO
	if err != nil {
		log.Fatalf("Failed map %v", err) // Jika terjadi error pada mapping, program akan berhenti dan menampilkan pesan error
	}
	res := service.userRepository.InsertUser(userToCreate) // Memanggil repository untuk membuat user baru
	return res                                             // Mengembalikan user yang telah dibuat
}

// FindByEmail adalah implementasi fungsi FindByEmail dari AuthService
func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email) // Memanggil repository untuk mencari user berdasarkan email
}

// IsDuplicateEmail adalah implementasi fungsi IsDuplicateEmail dari AuthService
func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email) // Memanggil repository untuk memeriksa apakah email sudah digunakan
	return !(res.Error == nil)                            // Mengembalikan true jika email sudah digunakan, false jika belum
}

// comparePassword adalah fungsi untuk membandingkan password yang dihash dengan password input
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)                                 // Mengkonversi password yang dihash menjadi byte slice
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword) // Membandingkan password yang dihash dengan password input
	if err != nil {
		return false // Jika tidak cocok, mengembalikan false
	}
	return true // Jika cocok, mengembalikan true
}
