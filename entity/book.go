package entity

// Book adalah model entitas yang merepresentasikan informasi buku dalam sistem
type Book struct {
	ID          uint64 `gorm:"primary_key:auto_increment"`                                                  // ID adalah identitas unik dari buku
	Title       string `gorm:"type:varchar(255)" json:"title"`                                              // Title adalah judul dari buku
	Description string `gorm:"type:text" json:"description"`                                                // Description adalah deskripsi atau konten dari buku
	UserID      uint64 `gorm:"not null" json:"-"`                                                           // UserID adalah ID dari user yang memiliki buku (disembunyikan dalam respons JSON)
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE, onDelete:CASCADE" json:"user"` // User adalah pemilik buku (relasi dengan model User)
}
