package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/dto"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/helper"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserController adalah interface yang mendefinisikan method-method yang dapat dipanggil untuk mengelola user
type UserController interface {
	Update(context *gin.Context)  // Method Update untuk meng-handle request update profile user
	Profile(context *gin.Context) // Method Profile untuk meng-handle request profile user
}

// userController adalah implementasi dari UserController
type userController struct {
	userService service.UserService // userService adalah service yang digunakan untuk operasi terkait user
	jwtService  service.JWTService  // jwtService adalah service yang digunakan untuk operasi terkait JWT
}

// NewUserController membuat instance baru dari UserController
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

// Update adalah method untuk meng-handle request update profile user
func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res) // Menampilkan response error bad request jika terjadi kesalahan pada DTO
		return
	}
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error()) // Panic jika terjadi kesalahan pada validasi token
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error()) // Panic jika terjadi kesalahan pada parsing user ID
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)    // Memanggil service untuk melakukan update user
	res := helper.BuildResponse(true, "OK!", u) // Membuat response sukses
	context.JSON(http.StatusOK, res)            // Mengirimkan response sukses
}

// Profile adalah method untuk meng-handle request profile user
func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error()) // Panic jika terjadi kesalahan pada validasi token
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)              // Memanggil service untuk mendapatkan profile user
	res := helper.BuildResponse(true, "OK!", user) // Membuat response sukses
	context.JSON(http.StatusOK, res)               // Mengirimkan response sukses
}
