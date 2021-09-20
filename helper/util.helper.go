package helper

import (
	"strconv"

	"github.com/dorman99/go_gin_mysql/entity"
	"github.com/gin-gonic/gin"
)

type paginationResponse struct {
	Results interface{} `json:"results"`
	Total   int         `json:"total"`
}

type BookResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func GenerateUsersPagination(data []entity.User) interface{} {
	return &paginationResponse{
		Results: data,
		Total:   len(data),
	}
}

func GenerateBooksPagination(data []BookResponse) interface{} {
	return paginationResponse{
		Results: data,
		Total:   len(data),
	}
}

func TransformBook(book entity.Book) BookResponse {
	b := BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.User.Name,
		Description: book.Description,
	}
	return b
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
