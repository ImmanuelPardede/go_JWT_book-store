package dto

// LoginDTO adalah model yang digunakan oleh client saat melakukan POST dari URL /login
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`       // Email adalah alamat email yang digunakan untuk login (wajib diisi)
	Password string `json:"password" form:"password" binding:"required"` // Password adalah kata sandi yang digunakan untuk login (wajib diisi)
}
