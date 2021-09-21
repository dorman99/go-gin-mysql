package service

import (
	"github.com/dorman99/go_gin_mysql/dto"
	"github.com/dorman99/go_gin_mysql/entity"
	repository "github.com/dorman99/go_gin_mysql/repo"
)

type BookService interface {
	Create(book dto.BookCreateDTO) entity.Book
	FindAll(limit uint64, skip uint64) []entity.Book
	FindByUser(userId uint64, limit uint64, skip uint64) []entity.Book
	Find(bookId uint64) entity.Book
	Update(book dto.BookUpdateDTO) entity.Book
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

func (s *bookService) FindByUser(userId uint64, limit uint64, skip uint64) []entity.Book {
	return s.bookRepo.FindByUser(userId, limit, skip)
}

func (s *bookService) Find(id uint64) entity.Book {
	return s.bookRepo.Find(id)
}

func (s *bookService) Update(book dto.BookUpdateDTO) entity.Book {
	return s.bookRepo.Update(book)
}
