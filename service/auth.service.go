package service

import (
	"log"

	common "github.com/dorman99/go_gin_mysql/common/service"
	"github.com/dorman99/go_gin_mysql/dto"
	"github.com/dorman99/go_gin_mysql/entity"
	repository "github.com/dorman99/go_gin_mysql/repo"
	"github.com/mashingan/smapping"
)

type AuthService interface {
	Register(user dto.RegisterDTO) entity.User
	Login(username string, password string) interface{}
}

type authService struct {
	userRepo repository.UserRepository
	bcrypt   common.BcryptMine
}

func NewAuthService(userRepo repository.UserRepository, bcrypt common.BcryptMine) AuthService {
	return &authService{
		userRepo: userRepo,
		bcrypt:   bcrypt,
	}
}

func (service *authService) Register(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	x := smapping.MapFields(&user)

	err := smapping.FillStruct(&userToCreate, x)
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	userToCreate.Password, err = service.bcrypt.Hash(userToCreate.Password)
	if err != nil {
		panic(err)
	}
	res := service.userRepo.Insert(userToCreate)
	return res
}

func (service *authService) Login(username string, password string) interface{} {
	res := service.userRepo.FindByUsername(username)
	if v, ok := res.(entity.User); ok {
		comparedPassword := service.bcrypt.Compare(password, v.Password)
		if comparedPassword {
			return res
		}
		return false
	}
	return false
}
