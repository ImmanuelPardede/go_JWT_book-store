package service

import (
	"log" // Mengimport package log untuk logging

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/dto"        // Mengimport package dto untuk DTO (Data Transfer Object)
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/entity"     // Mengimport package entity untuk model entitas
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/repository" // Mengimport package repository untuk interaksi dengan database
	"github.com/mashingan/smapping"                             // Mengimport package smapping untuk mapping struct
)

// UserService adalah interface yang mendefinisikan fungsi yang harus diimplementasikan oleh service user
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User // Fungsi untuk mengupdate user
	Profile(userID string) entity.User         // Fungsi untuk mendapatkan profil user berdasarkan ID
}

// userService adalah implementasi dari UserService
type userService struct {
	userRepository repository.UserRepository // Menggunakan repository untuk interaksi dengan database user
}

// NewUserService adalah constructor untuk userService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// Update adalah implementasi fungsi Update dari UserService
func (service *userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}                                        // Mendeklarasikan variabel untuk menyimpan data user yang akan diupdate
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user)) // Mengisi struct userToUpdate dengan data dari DTO
	if err != nil {
		log.Fatalf("Failed map %v:", err) // Jika terjadi error pada mapping, program akan berhenti dan menampilkan pesan error
	}
	updateUser := service.userRepository.UpdateUser(userToUpdate) // Memanggil repository untuk melakukan update data user
	return updateUser                                             // Mengembalikan data user yang telah diupdate
}

// Profile adalah implementasi fungsi Profile dari UserService
func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID) // Memanggil repository untuk mendapatkan profil user berdasarkan ID
}
