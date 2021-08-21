package service

import (
	repository "github.com/dorman99/go_gin_mysql/repo"
)

type UserService interface {
	FindByUsername(username string) interface{}
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service *userService) FindByUsername(username string) interface{} {
	return service.userRepo.FindByUsername(username)
}
