package service

import (
	"github.com/dorman99/go_gin_mysql/entity"
	repository "github.com/dorman99/go_gin_mysql/repo"
)

type UserService interface {
	FindByUsername(username string) interface{}
	Find(id uint64) entity.User
	FindAll(limit uint64, skip uint64) []entity.User
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

func (service *userService) Find(id uint64) entity.User {
	return service.userRepo.Find(id)
}

func (service *userService) FindAll(limit uint64, skip uint64) []entity.User {
	return service.userRepo.FindAll(limit, skip)
}

func (service *userService) Update(user entity.User) entity.User {
	return service.userRepo.Update(user)
}

func (service *userService) Remove(id uint64) entity.User {
	return service.userRepo.Remove(id)
}
