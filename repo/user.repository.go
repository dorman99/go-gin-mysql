package repository

import (
	"github.com/dorman99/go_gin_mysql/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(id uint64) entity.User
	Insert(user entity.User) entity.User
	Update(user entity.User) entity.User
	FindByUsername(username string) interface{}
	Remove(id uint64) entity.User
	FindAll(limit uint64, skip uint64) []entity.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Find(UserId uint64) entity.User {
	var user entity.User
	db.connection.Find(&user, UserId)
	return user
}

func (db *userConnection) FindAll(limit uint64, skip uint64) []entity.User {
	var users []entity.User
	db.connection.Where("deleted = ?", false).Limit(int(limit)).Offset(int(skip)).Find(&users)
	return users
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
		return res.Error
	}
	return user
}

func (db *userConnection) Remove(id uint64) entity.User {
	var user entity.User
	res := db.connection.Where("id", id).Update("deleted", true).Take(&user)
	if res.Error != nil {
		panic(res.Error)
	}
	return user
}
