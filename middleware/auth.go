package middleware

import (
	"log"
	"net/http"

	"github.com/dorman99/go_gin_mysql/common/server"
	common "github.com/dorman99/go_gin_mysql/common/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthJWt(jwtServ common.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		println(authHeader)
		if authHeader == "" {
			response := server.BuildErrorResponse("Invalid Header", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtServ.VerifyToken(authHeader)
		if err != nil {
			response := server.BuildErrorResponse("Invalid Token", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("UserId: ", claims["userId"])
			log.Println("Issuer: ", claims["issuer"])
		} else {
			response := server.BuildErrorResponse("Invalid Token", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
	}
}
