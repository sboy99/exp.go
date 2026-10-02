package repositories

import "github.com/sboy99/exp.go/go-book-store/internal/domain/models"

type IBookRepository interface {
	Create(name string, price float64) *models.Book
	List() *[]models.Book
	GetOne(id int64) *models.Book
}
