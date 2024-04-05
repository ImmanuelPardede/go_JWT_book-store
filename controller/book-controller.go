package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/dto"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/entity"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/helper"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// BookController adalah interface yang mendefinisikan method-method yang dapat dipanggil untuk mengelola buku
type BookController interface {
	All(context *gin.Context)      // Method All untuk meng-handle request mendapatkan semua buku
	FindByID(context *gin.Context) // Method FindByID untuk meng-handle request mendapatkan buku berdasarkan ID
	Insert(context *gin.Context)   // Method Insert untuk meng-handle request menambahkan buku baru
	Update(context *gin.Context)   // Method Update untuk meng-handle request update buku
	Delete(context *gin.Context)   // Method Delete untuk meng-handle request menghapus buku
}

// bookController adalah implementasi dari BookController
type bookController struct {
	bookService service.BookService // bookService adalah service yang digunakan untuk operasi terkait buku
	jwtService  service.JWTService  // jwtService adalah service yang digunakan untuk operasi terkait JWT
}

// NewBookController membuat instance baru dari BookController
func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
	return &bookController{
		bookService: bookServ,
		jwtService:  jwtServ,
	}
}

// All adalah method untuk meng-handle request mendapatkan semua buku
func (c *bookController) All(context *gin.Context) {
	var books []entity.Book = c.bookService.All()  // Mendapatkan semua buku dari service
	res := helper.BuildResponse(true, "OK", books) // Membuat response sukses
	context.JSON(http.StatusOK, res)               // Mengirimkan response sukses
}

// FindByID adalah method untuk meng-handle request mendapatkan buku berdasarkan ID
func (c *bookController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0) // Mendapatkan ID buku dari request
	if err != nil {
		res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{}) // Menampilkan response error jika tidak ada ID
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
	}
	var book entity.Book = c.bookService.FIndById(id) // Mendapatkan buku berdasarkan ID dari service
	if (book == entity.Book{}) {                      // Jika buku tidak ditemukan
		res := helper.BuildErrorResponse("Data not found", "No Data with given id", helper.EmptyObj{}) // Menampilkan response error data tidak ditemukan
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", book) // Membuat response sukses
		context.JSON(http.StatusOK, res)              // Mengirimkan response sukses
	}
}

// Insert adalah method untuk meng-handle request menambahkan buku baru
func (c *bookController) Insert(context *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := context.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{}) // Menampilkan response error jika terjadi kesalahan pada DTO
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader) // Mendapatkan ID user dari JWT token
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			bookCreateDTO.UserID = convertedUserID
		}
		result := c.bookService.Insert(bookCreateDTO)        // Menambahkan buku baru melalui service
		response := helper.BuildResponse(true, "OK", result) // Membuat response sukses
		context.JSON(http.StatusOK, response)                // Mengirimkan response sukses
	}
}

// Update adalah method untuk meng-handle request update buku
func (c *bookController) Update(context *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := context.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{}) // Menampilkan response error jika terjadi kesalahan pada DTO
		context.JSON(http.StatusBadRequest, res)
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic((errToken.Error())) // Panic jika terjadi kesalahan pada validasi token
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, bookUpdateDTO.ID) { // Memeriksa apakah user memiliki izin untuk mengedit buku
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			bookUpdateDTO.UserID = id
		}
		result := c.bookService.Update(bookUpdateDTO)        // Mengupdate buku melalui service
		response := helper.BuildResponse(true, "OK", result) // Membuat response sukses
		context.JSON(http.StatusOK, response)                // Mengirimkan response sukses
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{}) // Menampilkan response error jika user tidak memiliki izin
		context.JSON(http.StatusForbidden, response)
	}

}

// Delete adalah method untuk meng-handle request menghapus buku
func (c *bookController) Delete(context *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(context.Param("id"), 0, 0) // Mendapatkan ID buku dari request
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObj{}) // Menampilkan response error jika tidak ada ID
		context.JSON(http.StatusBadRequest, response)
	}
	book.ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic((errToken.Error())) // Panic jika terjadi kesalahan pada validasi token
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.bookService.IsAllowedToEdit(userID, book.ID) { // Memeriksa apakah user memiliki izin untuk mengedit buku
		c.bookService.Delete(book)                                     // Menghapus buku melalui service
		res := helper.BuildResponse(true, "Delete", helper.EmptyObj{}) // Membuat response sukses
		context.JSON(http.StatusOK, res)                               // Mengirimkan response sukses
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{}) // Menampilkan response error jika user tidak memiliki izin
		context.JSON(http.StatusForbidden, response)
	}
}

// getUserIDByToken adalah method untuk mendapatkan ID user dari JWT token
func (c *bookController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
