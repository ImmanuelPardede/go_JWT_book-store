package repository

import (
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/entity" // Mengimport package entity untuk model entitas
	"gorm.io/gorm"                                          // Mengimport package gorm untuk ORM
)

// BookRepository adalah interface yang mendefinisikan fungsi yang harus diimplementasikan oleh repository Book
type BookRepository interface {
	InsertBook(b entity.Book) entity.Book // Fungsi untuk menyimpan buku baru
	UpdateBook(b entity.Book) entity.Book // Fungsi untuk mengupdate buku
	DeleteBook(b entity.Book)             // Fungsi untuk menghapus buku
	AllBook() []entity.Book               // Fungsi untuk mendapatkan semua buku
	FindBookID(bookID uint64) entity.Book // Fungsi untuk mencari buku berdasarkan ID
}

// bookConnection adalah implementasi dari BookRepository
type bookConnection struct {
	connection *gorm.DB // Koneksi database menggunakan gorm
}

// NewBookRepository adalah constructor untuk bookConnection
func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConn,
	}
}

// InsertBook adalah implementasi fungsi InsertBook dari BookRepository
func (db *bookConnection) InsertBook(b entity.Book) entity.Book {
	db.connection.Save(&b)                 // Menyimpan buku ke database
	db.connection.Preload("User").Find(&b) // Mengambil buku yang baru disimpan dengan relasi User
	return b                               // Mengembalikan buku yang telah disimpan
}

// UpdateBook adalah implementasi fungsi UpdateBook dari BookRepository
func (db *bookConnection) UpdateBook(b entity.Book) entity.Book {
	db.connection.Save(&b)                 // Mengupdate buku ke database
	db.connection.Preload("User").Find(&b) // Mengambil buku yang telah diupdate dengan relasi User
	return b                               // Mengembalikan buku yang telah diupdate
}

// DeleteBook adalah implementasi fungsi DeleteBook dari BookRepository
func (db *bookConnection) DeleteBook(b entity.Book) {
	db.connection.Delete(&b) // Menghapus buku dari database
}

// FindBookID adalah implementasi fungsi FindBookID dari BookRepository
func (db *bookConnection) FindBookID(bookID uint64) entity.Book {
	var book entity.Book
	db.connection.Preload("User").Find(&book, bookID) // Mengambil buku berdasarkan ID dengan relasi User
	return book                                       // Mengembalikan buku yang ditemukan
}

// AllBook adalah implementasi fungsi AllBook dari BookRepository
func (db *bookConnection) AllBook() []entity.Book {
	var books []entity.Book
	db.connection.Preload("User").Find(&books) // Mengambil semua buku dengan relasi User
	return books                               // Mengembalikan semua buku yang ditemukan
}
