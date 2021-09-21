package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dorman99/go_gin_mysql/common/server"
	"github.com/dorman99/go_gin_mysql/dto"
	"github.com/dorman99/go_gin_mysql/entity"
	"github.com/dorman99/go_gin_mysql/helper"
	"github.com/dorman99/go_gin_mysql/service"
	"github.com/gin-gonic/gin"
)

type BookController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByUser(ctx *gin.Context)
	UpdateSelf(ctx *gin.Context)
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

func (con *bookController) FindByUser(ctx *gin.Context) {
	var userHead server.HeaderRequest
	headReq, _ := ctx.Get("user")
	// ubah ke bytes lalu di marshal
	// baca ini https://idineshkrishnan.com/json-marshalling-and-unmarshalling-in-golang/
	bytes, errM := json.Marshal(headReq)
	if errM != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errM)
		return
	}
	err := json.Unmarshal(bytes, &userHead)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	userId, err_uin := strconv.ParseInt(userHead.UserId, 10, 64)
	if err_uin != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err_uin)
		return
	}

	limit, skip := helper.GetLimitSkip(ctx)
	books := con.bookService.FindByUser(uint64(userId), limit, skip)
	var bookList []helper.BookResponse
	for _, v := range books {
		b := helper.TransformBook(v)
		bookList = append(bookList, b)
	}
	list := helper.GenerateBooksPagination(bookList)
	response := server.BuildResponse(true, "success", list)
	ctx.JSON(http.StatusOK, response)
}

func (con *bookController) UpdateSelf(ctx *gin.Context) {
	var userRequestHeader server.HeaderRequest
	headedReq, _ := ctx.Get("user")

	bytes, errM := json.Marshal(headedReq)
	if errM != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errM)
		return
	}

	errUm := json.Unmarshal(bytes, &userRequestHeader)
	if errUm != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errUm)
		return
	}

	userId, errP := strconv.ParseInt(userRequestHeader.UserId, 0, 0)
	if errP != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errP)
		return
	}

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var bookUpdateDto = dto.BookUpdateDTO{
		ID:     uint64(id),
		UserID: userId,
	}

	errDto := ctx.ShouldBind(&bookUpdateDto)
	if errDto != nil {
		response := server.BuildErrorResponse("Bad Request", errDto.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	book := con.bookService.Find(uint64(id))
	if (book == entity.Book{}) {
		response := server.BuildErrorResponse("not found", "book is not found", nil)
		ctx.JSON(http.StatusNotFound, response)
		return
	} else if book.UserID != uint64(userId) {
		response := server.BuildErrorResponse("forbidden", "its not yours", nil)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	update := con.bookService.Update(bookUpdateDto)
	bu := helper.TransformBook(update)
	response := server.BuildResponse(true, "success", bu)
	ctx.JSON(http.StatusOK, response)
}
