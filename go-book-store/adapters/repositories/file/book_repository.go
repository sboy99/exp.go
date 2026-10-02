package file

import (
	"encoding/json"

	"github.com/sboy99/exp.go/go-book-store/infra/exceptions"
	"github.com/sboy99/exp.go/go-book-store/internal/domain/models"
	"github.com/sboy99/exp.go/go-book-store/internal/helpers"
)

type BookFileRepository struct{}

func (r *BookFileRepository) Create(name string, price float64) *models.Book {
	author := models.Author{
		ID:   helpers.GenRandomId(),
		Name: "Mr Smith",
		Sex:  "Male",
	}

	book := models.Book{
		ID:     helpers.GenRandomId(),
		Name:   name,
		Price:  price,
		Author: &author,
	}

	// add books to array
	books = append(books, book)
	booksJSON, err := json.Marshal(books)
	exceptions.HandleJsonMarshalException(err)

	// save book to database
	helpers.SaveFile("books.json", booksJSON)

	return &book
}

func (r *BookFileRepository) List() *[]models.Book {
	return &books
}

func (r *BookFileRepository) GetOne(id int64) *models.Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}

	return nil
}
