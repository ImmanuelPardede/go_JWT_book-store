package middleware

import (
	"log"      // Mengimport package log untuk logging
	"net/http" // Mengimport package http untuk handling HTTP requests

	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/helper"  // Mengimport package helper untuk fungsi bantuan
	"github.com/ImmanuelPardede/golang_gin_gorm_GWT/service" // Mengimport package service untuk JWTService
	"github.com/dgrijalva/jwt-go"                            // Mengimport package jwt-go untuk JWT (JSON Web Token)
	"github.com/gin-gonic/gin"                               // Mengimport package gin untuk web framework
)

// AuthorizeJWT adalah middleware untuk validasi token JWT yang diberikan oleh user
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization") // Mengambil header Authorization dari request
		if authHeader == "" {                      // Jika header Authorization tidak ditemukan
			response := helper.BuildErrorResponse("Failed to process request", "no token found", nil) // Membangun respons error
			c.AbortWithStatusJSON(http.StatusBadRequest, response)                                    // Mengirim respons error dengan status 400 Bad Request
			return                                                                                    // Menghentikan eksekusi middleware
		}

		token, err := jwtService.ValidateToken(authHeader) // Validasi token JWT
		if token.Valid {                                   // Jika token valid
			claims := token.Claims.(jwt.MapClaims)              // Mengambil claims dari token
			log.Println("claim[user_id] : ", claims["user_id"]) // Menampilkan user_id dari claims
			log.Println("Claim[issuer] :", claims["issuer"])    // Menampilkan issuer dari claims
		} else { // Jika token tidak valid
			log.Println(err)                                                              // Log pesan error
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil) // Membangun respons error
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)                      // Mengirim respons error dengan status 401 Unauthorized
		}
	}
}
