package route

import (
	"net/http"

	common "github.com/dorman99/go_gin_mysql/common/service"
	"github.com/dorman99/go_gin_mysql/controller"
	repository "github.com/dorman99/go_gin_mysql/repo"
	"github.com/dorman99/go_gin_mysql/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func V1Routes(r *gin.Engine, db *gorm.DB) *gin.RouterGroup {

	jwtService := common.NewJWTServ()
	bcryptService := common.NewBcrypt()

	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, bcryptService)

	authCon := controller.NewAuthController(userService, authService, jwtService)
	routesGroups := r.Group("/api/v1")
	{
		routesGroups.Handle(http.MethodPost, "/register", authCon.Register)
		routesGroups.Handle(http.MethodPost, "/login", authCon.Login)
	}

	return routesGroups

}
