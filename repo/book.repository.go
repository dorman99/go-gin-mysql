package repository

import (
	"github.com/dorman99/go_gin_mysql/dto"
	"github.com/dorman99/go_gin_mysql/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	Find(id uint64) entity.Book
	Create(book dto.BookCreateDTO) entity.Book
	Update(book dto.BookUpdateDTO) entity.Book
	FindByUser(id uint64, limit uint64, skip uint64) []entity.Book
	Remove(id uint64) entity.Book
	FindAll(limit uint64, skip uint64) []entity.Book
}

type bookConnection struct {
	connection *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookConnection{
		connection: db,
	}
}

func (db *bookConnection) FindAll(limit uint64, skip uint64) []entity.Book {
	var books []entity.Book
	res := db.connection.Where("deleted = ?", false).Limit(int(limit)).Offset(int(skip)).Preload("User").Find(&books)
	if res.Error != nil {
		panic(res.Error)
	}
	return books
}

func (db *bookConnection) Create(book dto.BookCreateDTO) entity.Book {
	bookCreated := entity.Book{
		Title:       book.Title,
		Description: book.Description,
		UserID:      uint64(book.UserID),
	}
	ress := db.connection.Save(&bookCreated)
	if ress.Error != nil {
		panic(ress.Error)
	}
	return bookCreated
}

func (db *bookConnection) Update(book dto.BookUpdateDTO) entity.Book {
	b := entity.Book{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		UserID:      uint64(book.UserID),
	}
	db.connection.Preload("User").Save(b).Take(&b)
	return b
}

func (db *bookConnection) Remove(id uint64) entity.Book {
	var book entity.Book
	res := db.connection.Where("id", id).Update("deleted", true).Take(&book)
	if res.Error != nil {
		panic(res.Error)
	}
	return book
}

func (db *bookConnection) Find(id uint64) entity.Book {
	var book entity.Book
	res := db.connection.Find(&book, id)
	if res.Error != nil {
		panic(res.Error)
	}
	return book
}

func (db *bookConnection) FindByUser(id uint64, limit uint64, skip uint64) []entity.Book {
	var books []entity.Book
	res := db.connection.Preload("User").Where("user_id = ? AND deleted = ?", id, false).Limit(int(limit)).Offset(int(skip)).Find(&books)
	if res.Error != nil {
		panic(res.Error)
	}
	return books
}
