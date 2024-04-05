package dto

// RegisterDTO digunakan saat client melakukan post dari URL /register
type RegisterDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`         // Name adalah nama lengkap user yang akan diregistrasi (wajib diisi)
	Email    string `json:"email" form:"email" binding:"required,email"` // Email adalah alamat email user yang akan diregistrasi (wajib diisi dan harus sesuai format email)
	Password string `json:"password" form:"password" binding:"required"` // Password adalah kata sandi user yang akan diregistrasi (wajib diisi)
}
