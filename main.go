package main

import (
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/config"     // Mengimport konfigurasi aplikasi
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/controller" // Mengimport controller aplikasi
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/middleware" // Mengimport middleware aplikasi
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/repository" // Mengimport repository aplikasi
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/service"    // Mengimport service aplikasi
	"github.com/gin-gonic/gin"                                  // Mengimport framework gin untuk routing HTTP
	"gorm.io/gorm"                                              // Mengimport ORM GORM untuk manipulasi database
)

// Inisialisasi variabel yang digunakan dalam aplikasi
var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()                      // Membuat koneksi database
	userRepository repository.UserRepository = repository.NewUserRepository(db)                      // Membuat repository user
	bookRepository repository.BookRepository = repository.NewBookRepository(db)                      // Membuat repository buku
	jwtService     service.JWTService        = service.NewJWTService()                               // Membuat service JWT
	userService    service.UserService       = service.NewUserService(userRepository)                // Membuat service user
	bookService    service.BookService       = service.NewBookService(bookRepository)                // Membuat service buku
	authService    service.AuthService       = service.NewAuthService(userRepository)                // Membuat service auth
	authController controller.AuthController = controller.NewAuthController(authService, jwtService) // Membuat controller auth
	userController controller.UserController = controller.NewUserController(userService, jwtService) // Membuat controller user
	bookController controller.BookController = controller.NewBookController(bookService, jwtService) // Membuat controller buku
)

// Fungsi utama aplikasi
func main() {
	defer config.CloseDatabaseConnection(db) // Menutup koneksi database secara defer
	r := gin.Default()                       // Menggunakan router default dari Gin

	authRoutes := r.Group("api/auth") // Membuat grup endpoint untuk auth
	{
		authRoutes.POST("/login", authController.Login)       // Endpoint login
		authRoutes.POST("/register", authController.Register) // Endpoint register
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService)) // Membuat grup endpoint untuk user dengan middleware JWT
	{
		userRoutes.GET("/profile", userController.Profile) // Endpoint profil user
		userRoutes.PUT("/profile", userController.Update)  // Endpoint update profil user
	}

	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService)) // Membuat grup endpoint untuk buku dengan middleware JWT
	{
		bookRoutes.GET("/", bookController.All)          // Endpoint untuk mendapatkan semua buku
		bookRoutes.POST("/", bookController.Insert)      // Endpoint untuk menyimpan buku baru
		bookRoutes.GET("/:id", bookController.FindByID)  // Endpoint untuk mencari buku berdasarkan ID
		bookRoutes.PUT("/:id", bookController.Update)    // Endpoint untuk mengupdate buku berdasarkan ID
		bookRoutes.DELETE("/:id", bookController.Delete) // Endpoint untuk menghapus buku berdasarkan ID
	}

	r.Run(":9090") // Menjalankan server pada port 9090
}
