package file

import (
	"encoding/json"

	"github.com/sboy99/exp.go/go-book-store/infra/exceptions"
	"github.com/sboy99/exp.go/go-book-store/internal/domain/models"
	"github.com/sboy99/exp.go/go-book-store/internal/helpers"
)

var books []models.Book

func LoadBooks() {
	booksData := helpers.LoadFile("books.json")
	err := json.Unmarshal(booksData, &books)
	exceptions.HandleJsonUnMarshalException(err)
}
