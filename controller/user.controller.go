package controller

import (
	"net/http"
	"strconv"

	"github.com/dorman99/go_gin_mysql/common/server"
	"github.com/dorman99/go_gin_mysql/entity"
	"github.com/dorman99/go_gin_mysql/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Find(context *gin.Context)
	FindAll(contet *gin.Context)
	// Update(id int, updateDTO dto.UserUpdateDTO) entity.User
	// Remove(id int)
}
type userController struct {
	userService service.UserService
}

func NewUserContoller(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) Find(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := server.BuildErrorResponse("Invalid Params", "Must Be Provided", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	var acc entity.User = c.userService.Find(id)
	if (acc == entity.User{}) {
		response := server.BuildErrorResponse("Invalid Id", "Not Found", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := server.BuildResponse(true, "Ok", acc)
	context.JSON(http.StatusOK, response)
}

func (c *userController) FindAll(context *gin.Context) {
	limit := uint64(10)
	skip := uint64(0)
	users := c.userService.FindAll(limit, skip)
	response := server.BuildResponse(true, "Ok", users)
	context.JSON(http.StatusOK, response)
}

func (c *userController) Update(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := server.BuildErrorResponse("Invalid Params", "Must Be Provided", nil)
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
	}

	user := c.userService.Find(id)
	if (user == entity.User{}) {
		response := server.BuildErrorResponse("Invalid Id", "Not Found", nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

}
