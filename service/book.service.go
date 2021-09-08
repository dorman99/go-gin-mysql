package service

import (
	"github.com/dorman99/go_gin_mysql/dto"
	"github.com/dorman99/go_gin_mysql/entity"
	repository "github.com/dorman99/go_gin_mysql/repo"
)

type BookService interface {
	Create(book dto.BookCreateDTO) entity.Book
	FindAll(limit uint64, skip uint64) []entity.Book
}

type bookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) Create(book dto.BookCreateDTO) entity.Book {
	return s.bookRepo.Create(book)
}

func (s *bookService) FindAll(limit uint64, skip uint64) []entity.Book {
	return s.bookRepo.FindAll(limit, skip)
}
