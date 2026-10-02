package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-book-store/adapters/controllers/dtos"
	"github.com/sboy99/exp.go/go-book-store/internal/domain/models"
	"github.com/sboy99/exp.go/go-book-store/internal/helpers"
	"github.com/sboy99/exp.go/go-book-store/internal/services"
)

type BookController struct {
	BookService services.IBookService
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var createBookDto dtos.CreateBookReq

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.Decode(&createBookDto)

	fmt.Println(createBookDto)

	// create the book
	book := c.BookService.CreateBook(createBookDto.Name, createBookDto.Price)
	ctx.JSON(http.StatusOK, *book)
}

func (c *BookController) ListBooks(ctx *gin.Context) {
	books := c.BookService.ListBooks()
	ctx.JSON(http.StatusOK, *books)
}

func (c *BookController) GetOneBook(ctx *gin.Context) {
	id := ctx.Param("id")
	bookID := helpers.ParseInt64(id)

	book := c.BookService.GetOneBook(bookID)
	ctx.JSON(http.StatusOK, *book)
}

func (s *BookController) UpdateOneBook(ctx *gin.Context) {
	var book models.Book
	ctx.JSON(http.StatusOK, book)
}

func (s *BookController) DeleteOneBook(ctx *gin.Context) {
	var book models.Book
	// todo: call remove service
	ctx.JSON(http.StatusOK, book)
}
