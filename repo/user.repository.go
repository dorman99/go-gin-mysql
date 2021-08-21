package repository

import (
	"github.com/dorman99/go_gin_mysql/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(id int64) entity.User
	Insert(user entity.User) entity.User
	Update(user entity.User) entity.User
	FindByUsername(username string) interface{}
	Remove(id int64) entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Find(UserId int64) entity.User {
	var user entity.User
	db.connection.Find(&user, UserId)
	return user
}

func (db *userConnection) Insert(user entity.User) entity.User {
	db.connection.Save(&user)
	return user
}

func (db *userConnection) Update(user entity.User) entity.User {
	db.connection.Save(&user)
	return user
}

func (db *userConnection) FindByUsername(username string) interface{} {
	var user entity.User
	res := db.connection.Where("username = ?", username).Take(&user)
	if res.Error != nil {
		return nil
	}
	return user
}

func (db *userConnection) Remove(id int64) entity.User {
	var user entity.User
	return user
}
