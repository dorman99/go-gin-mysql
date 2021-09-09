package middleware

import (
	"net/http"

	"github.com/dorman99/go_gin_mysql/common/server"
	common "github.com/dorman99/go_gin_mysql/common/service"
	"github.com/gin-gonic/gin"
)

func AuthJWt(jwtServ common.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := server.BuildErrorResponse("Invalid Header", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claims, err := jwtServ.VerifyToken(authHeader)
		if err != nil {
			response := server.BuildErrorResponse("Invalid Token", "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("user", claims)
	}
}
