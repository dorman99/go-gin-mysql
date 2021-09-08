package helper

import "github.com/dorman99/go_gin_mysql/entity"

type paginationResponse struct {
	data  interface{}
	total int
}

func GenerateUsersPagination(data []entity.User) interface{} {
	return &paginationResponse{
		data:  data,
		total: len(data),
	}
}
