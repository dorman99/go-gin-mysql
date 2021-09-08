package controller

import (
	"net/http"
	"strconv"

	"github.com/dorman99/go_gin_mysql/common/server"
	"github.com/dorman99/go_gin_mysql/dto"
	"github.com/dorman99/go_gin_mysql/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &bookController{
		bookService: bookService,
	}
}

func (con *bookController) Create(ctx *gin.Context) {
	var bookDto dto.BookCreateDTO
	errDto := ctx.ShouldBind(&bookDto)
	if errDto != nil {
		response := server.BuildErrorResponse("Bad Request", errDto.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	book := con.bookService.Create(bookDto)
	response := server.BuildResponse(true, "Success", book)
	ctx.JSON(http.StatusOK, response)
}

func (con *bookController) FindAll(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	skip := ctx.DefaultQuery("skip", "0")
	_limit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		panic(err)
	}
	_skip, err := strconv.ParseInt(skip, 10, 64)
	if err != nil {
		panic(err)
	}
	books := con.bookService.FindAll(uint64(_limit), uint64(_skip))
	response := server.BuildResponse(true, "Success", books)
	ctx.JSON(http.StatusOK, response)
}