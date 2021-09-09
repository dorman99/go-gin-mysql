package route

import (
	"net/http"

	common "github.com/dorman99/go_gin_mysql/common/service"
	"github.com/dorman99/go_gin_mysql/controller"
	"github.com/dorman99/go_gin_mysql/middleware"
	repository "github.com/dorman99/go_gin_mysql/repo"
	"github.com/dorman99/go_gin_mysql/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func V1Routes(r *gin.Engine, db *gorm.DB) *gin.RouterGroup {

	jwtService := common.NewJWTServ()
	bcryptService := common.NewBcrypt()

	authMiddleware := middleware.AuthJWt(jwtService)

	userRepo := repository.NewUserRepository(db)
	bookRepo := repository.NewBookRepository(db)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, bcryptService)
	bookService := service.NewBookService(bookRepo)

	authCon := controller.NewAuthController(userService, authService, jwtService)
	userCon := controller.NewUserContoller(userService)
	bookCon := controller.NewBookController(bookService)
	routesGroups := r.Group("/api/v1")
	{
		routesGroups.Handle(http.MethodPost, "/auth/register", authCon.Register)
		routesGroups.Handle(http.MethodPost, "/auth/login", authCon.Login)
		routesGroups.Use(authMiddleware)
		routesGroups.Handle(http.MethodGet, "/users/:id", userCon.Find)
		routesGroups.Handle(http.MethodGet, "/users", userCon.FindAll)
		routesGroups.Handle(http.MethodPost, "/books", bookCon.Create)
		routesGroups.Handle(http.MethodGet, "/books", bookCon.FindAll)
		routesGroups.Handle(http.MethodGet, "/books/users/self", bookCon.FindByUser)
	}

	return routesGroups

}
