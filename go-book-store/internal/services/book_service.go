package services

import (
	"github.com/sboy99/exp.go/go-book-store/internal/domain/models"
	"github.com/sboy99/exp.go/go-book-store/internal/repositories"
)

type BookService struct {
	BookRepo repositories.IBookRepository
}

func (s *BookService) CreateBook(name string, price float64) *models.Book {
	book := s.BookRepo.Create(name, price)
	return book
}

func (s *BookService) ListBooks() *[]models.Book {
	books := s.BookRepo.List()
	return books
}

func (s *BookService) GetOneBook(id int64) *models.Book {
	book := s.BookRepo.GetOne(id)
	return book
}
