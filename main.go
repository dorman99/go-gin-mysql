package main

import (
	common "github.com/dorman99/go_gin_mysql/common/service"
	"github.com/dorman99/go_gin_mysql/config"
	"github.com/dorman99/go_gin_mysql/controller"
	"github.com/dorman99/go_gin_mysql/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDBConnection(db)
	r := gin.Default()
	authRoutes := r.Group("/api/auth")
	authRoutes.Use(middleware.AuthJWt(common.NewJWTSer()))
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
