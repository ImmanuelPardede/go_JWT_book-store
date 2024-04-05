package dto

// BookUpdateDTO adalah model yang digunakan oleh client saat melakukan update pada buku
type BookUpdateDTO struct {
	ID          uint64 `json:"id" binding:"required"`                             // ID adalah identitas unik dari buku yang akan diupdate (wajib diisi)
	Title       string `json:"title" form:"title" binding:"required"`             // Title adalah judul baru dari buku yang akan diupdate (wajib diisi)
	Description string `json:"description" form:"description" binding:"required"` // Description adalah deskripsi baru dari buku yang akan diupdate (wajib diisi)
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`        // UserID adalah ID user yang memiliki buku (opsional, bisa kosong)
}

// BookCreateDTO adalah model yang digunakan oleh client saat membuat buku baru
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`             // Title adalah judul dari buku yang akan dibuat (wajib diisi)
	Description string `json:"description" form:"description" binding:"required"` // Description adalah deskripsi dari buku yang akan dibuat (wajib diisi)
	UserID      uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`        // UserID adalah ID user yang membuat buku (opsional, bisa kosong)
}
