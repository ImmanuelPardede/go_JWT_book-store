package helper

import (
	"strings" // Mengimport package strings untuk manipulasi string
)

// Response digunakan untuk bentuk statis dari respons JSON
type Response struct {
	Status  bool        `json:"status"`  // Status respons
	Message string      `json:"message"` // Pesan respons
	Errors  interface{} `json:"errors"`  // Error yang terjadi
	Data    interface{} `json:"data"`    // Data yang dikirimkan
}

// EmptyObj digunakan ketika data tidak ingin menjadi null pada JSON
type EmptyObj struct{}

// BuildResponse adalah metode untuk menyusun respons sukses yang dinamis
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,  // Mengisi status respons
		Message: message, // Mengisi pesan respons
		Errors:  nil,     // Mengosongkan daftar error karena respons sukses
		Data:    data,    // Mengisi data respons
	}
	return res // Mengembalikan respons yang telah dibuat
}

// BuildErrorResponse adalah metode untuk menyusun respons gagal yang dinamis
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splitedError := strings.Split(err, "\n") // Membagi pesan error menjadi beberapa baris jika ada newline
	res := Response{
		Status:  false,        // Status respons gagal
		Message: message,      // Pesan respons gagal
		Errors:  splitedError, // Mengisi daftar error dengan pesan error yang telah dibagi
		Data:    data,         // Mengisi data respons
	}
	return res // Mengembalikan respons yang telah dibuat
}
