package controller

import (
	"net/http"
	"strconv"

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/dto"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/entity"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/helper"
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/service"
	"github.com/gin-gonic/gin"
)

// AuthController interface adalah kontrak untuk controller ini
type AuthController interface {
	Login(ctx *gin.Context)    // Method untuk meng-handle request login
	Register(ctx *gin.Context) // Method untuk meng-handle request registrasi
}

// authController adalah implementasi dari AuthController
type authController struct {
	authService service.AuthService // authService adalah service yang digunakan untuk operasi terkait auth
	jwtService  service.JWTService  // jwtService adalah service yang digunakan untuk operasi terkait JWT
}

// NewAuthController membuat instance baru dari AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Login adalah method untuk meng-handle request login
func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{}) // Menampilkan response error jika terjadi kesalahan pada DTO
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password) // Verifikasi kredensial user
	if v, ok := authResult.(entity.User); ok {                                      // Jika kredensial valid
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10)) // Generate token JWT
		v.Token = generatedToken
		response := helper.BuildResponse(true, "OK!", v) // Membuat response sukses dengan token
		ctx.JSON(http.StatusOK, response)                // Mengirimkan response sukses
		return
	}
	response := helper.BuildErrorResponse("Please check again your credential", "Invalid Credential", helper.EmptyObj{}) // Menampilkan response error kredensial tidak valid
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Register adalah method untuk meng-handle request registrasi
func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{}) // Menampilkan response error jika terjadi kesalahan pada DTO
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if !c.authService.IsDuplicateEmail(registerDTO.Email) { // Memeriksa duplikasi email
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate email", helper.EmptyObj{}) // Menampilkan response error jika email sudah terdaftar
		ctx.JSON(http.StatusConflict, response)                                                                  // Mengirimkan response conflict
	} else {
		createdUser := c.authService.CreateUser(registerDTO)                        // Membuat user baru melalui service
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10)) // Generate token JWT
		createdUser.Token = token
		response := helper.BuildResponse(true, "OK!", createdUser) // Membuat response sukses dengan token
		ctx.JSON(http.StatusCreated, response)                     // Mengirimkan response sukses
	}
}
