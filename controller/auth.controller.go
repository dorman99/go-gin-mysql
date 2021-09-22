package controller

import (
	"net/http"

	"github.com/dorman99/go_gin_mysql/common/server"
	common "github.com/dorman99/go_gin_mysql/common/service"
	"github.com/dorman99/go_gin_mysql/dto"
	"github.com/dorman99/go_gin_mysql/entity"
	"github.com/dorman99/go_gin_mysql/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	userService service.UserService
	authService service.AuthService
	jwtService  common.JWTService
}

func NewAuthController(userService service.UserService, authService service.AuthService, jwtService common.JWTService) AuthController {
	return &authController{
		userService: userService,
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDto dto.LoginDTO
	errDto := ctx.ShouldBind(&loginDto)
	if errDto != nil {
		response := server.BuildErrorResponse("Failed To Parse", errDto.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	loginValidating := c.authService.Login(loginDto.Username, loginDto.Password)
	if v, ok := loginValidating.(entity.User); ok {
		user := dto.LoginResponseDto{
			ID:       v.ID,
			Name:     v.Name,
			Username: v.Username,
		}
		generateToken, refreshToken := c.jwtService.GeneratePairToken(v)
		user.Token = generateToken
		user.Refresh = refreshToken
		response := server.BuildResponse(true, "OK", user)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := server.BuildErrorResponse("Invalid Username / Password", "Invalid Credentials", nil)
	ctx.JSON(http.StatusBadRequest, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDto := ctx.ShouldBind(&registerDTO)
	if errDto != nil {
		response := server.BuildErrorResponse("Invalid Object", errDto.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user := c.userService.FindByUsername(registerDTO.Username)
	if _, ok := user.(entity.User); ok {
		response := server.BuildErrorResponse("Duplicate", "Duplicate Username", nil)
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	}

	regist := c.authService.Register(registerDTO)
	response := server.BuildResponse(true, "OK", regist)
	ctx.JSON(http.StatusOK, response)
}
