package service

import (
	"fmt"  // Mengimport package fmt untuk formatting dan printing
	"os"   // Mengimport package os untuk mengakses environment variable
	"time" // Mengimport package time untuk mengelola waktu

	"github.com/dgrijalva/jwt-go" // Mengimport package jwt-go untuk JWT (JSON Web Token)
)

// JWTService adalah interface yang mendefinisikan fungsi yang harus diimplementasikan oleh service JWT
type JWTService interface {
	GenerateToken(userID string) string             // Fungsi untuk generate token JWT
	ValidateToken(token string) (*jwt.Token, error) // Fungsi untuk validasi token JWT
}

// jwtCustomClaim adalah struct untuk menyimpan custom claim JWT
type jwtCustomClaim struct {
	UserID             string `json:"user_id"` // Field untuk user ID dalam claim JWT
	jwt.StandardClaims        // Field untuk standard claims JWT
}

// jwtService adalah implementasi dari JWTService
type jwtService struct {
	secretKey string // Kunci rahasia untuk signing JWT
	issuer    string // Issuer untuk JWT
}

// NewJWTService adalah constructor untuk jwtService
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ImmanuelPardede", // Set issuer JWT
		secretKey: getSecretKey(),    // Set kunci rahasia JWT
	}
}

// getSecretKey adalah fungsi untuk mendapatkan kunci rahasia JWT dari environment variable
func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET") // Mendapatkan nilai dari environment variable JWT_SECRET
	if secretKey == "" {                 // Jika nilai tidak ditemukan, gunakan default "ImmanuelPardede"
		secretKey = "ImmanuelPardede"
	}
	return secretKey
}

// GenerateToken adalah implementasi fungsi GenerateToken dari JWTService
func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{ // Membuat custom claim JWT
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(), // Token akan kedaluwarsa dalam 1 tahun
			Issuer:    j.issuer,                           // Mengatur issuer JWT
			IssuedAt:  time.Now().Unix(),                  // Waktu pembuatan token
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Membuat token JWT
	t, err := token.SignedString([]byte(j.secretKey))          // Mengesahkan token dengan menggunakan kunci rahasia
	if err != nil {
		panic(err) // Panic jika terjadi error dalam pembuatan token
	}
	return t // Mengembalikan token JWT yang telah dibuat
}

// ValidateToken adalah implementasi fungsi ValidateToken dari JWTService
func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"]) // Error jika metode signing tidak sesuai
		}
		return []byte(j.secretKey), nil // Mengembalikan kunci rahasia untuk validasi token
	})
}
