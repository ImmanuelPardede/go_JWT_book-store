package entity

// User adalah model entitas yang merepresentasikan data pengguna dalam sistem
type User struct {
	ID       uint64  `gorm:"primary_key:auto_increament" json:"id"`      // ID adalah identitas unik dari user
	Name     string  `gorm:"type:varchar(255)" json:"name"`              // Name adalah nama lengkap dari user
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"` // Email adalah alamat email dari user
	Password string  `gorm:"->;<-;not null" json:"-"`                    // Password adalah kata sandi dari user (disembunyikan dalam respons JSON)
	Token    string  `gorm:"-" json:"token,omitempty"`                   // Token adalah token JWT yang diterbitkan kepada user saat login (tidak disimpan dalam database)
	Books    *[]Book `json:"books,omitempty"`                            // Books adalah daftar buku yang dimiliki oleh user (opsional, bisa kosong)
}
