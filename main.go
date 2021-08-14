package main

import (
	"github.com/dorman99/go_gin_mysql/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	// authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDBConnection(db)
	r := gin.Default()
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login")
	}

	r.Run()
}
