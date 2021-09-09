package helper

import (
	"strconv"

	"github.com/dorman99/go_gin_mysql/entity"
	"github.com/gin-gonic/gin"
)

type paginationResponse struct {
	data  interface{}
	total int
}

type BookResponse struct {
	id          uint64
	title       string
	author      string
	description string
}

func GenerateUsersPagination(data []entity.User) interface{} {
	return &paginationResponse{
		data:  data,
		total: len(data),
	}
}

func GenerateBooksPagination(data []BookResponse) interface{} {
	return &paginationResponse{
		data:  data,
		total: len(data),
	}
}

func TransformBook(book entity.Book) BookResponse {
	return BookResponse{
		id:          book.ID,
		title:       book.Title,
		author:      book.User.Name,
		description: book.Description,
	}
}

func GetLimitSkip(c *gin.Context) (l uint64, s uint64) {
	limit := c.DefaultQuery("limit", "10")
	skip := c.DefaultQuery("skip", "0")
	_limit, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		panic(err)
	}
	_skip, err := strconv.ParseInt(skip, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint64(_limit), uint64(_skip)
}
