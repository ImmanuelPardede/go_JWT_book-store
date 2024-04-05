package dto

// UserUpdateDTO digunakan oleh client saat melakukan update profile menggunakan metode PUT
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`                                                    // ID adalah identitas unik dari user yang akan diupdate
	Name     string `json:"name" form:"name" binding:"required"`                             // Name adalah nama lengkap user yang akan diupdate (wajib diisi)
	Email    string `json:"email" form:"email" binding:"required,email"`                     // Email adalah alamat email user yang akan diupdate (wajib diisi dan harus sesuai format email)
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required"` // Password adalah kata sandi user yang akan diupdate (wajib diisi)
}

// UserCreateDTO digunakan oleh client saat membuat user baru (kode ini di-comment karena tidak digunakan saat ini)
// type UserCreateDTO struct {
//     Name     string `json:"name" form:"name" binding:"required"`  // Name adalah nama lengkap user yang akan dibuat (wajib diisi)
//     Email    string `json:"email" form:"email" binding:"required" validate:"email"`  // Email adalah alamat email user yang akan dibuat (wajib diisi dan harus sesuai format email)
//     Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`  // Password adalah kata sandi user yang akan dibuat (wajib diisi dan minimal 6 karakter)
// }
