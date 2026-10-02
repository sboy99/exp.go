package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-book-store/adapters/controllers"
	"github.com/sboy99/exp.go/go-book-store/adapters/repositories/file"
	"github.com/sboy99/exp.go/go-book-store/internal/services"
)

func RegisterBookRoutes(router *gin.Engine) {
	bookFileRepository := &file.BookFileRepository{}
	bookService := &services.BookService{
		BookRepo: bookFileRepository,
	}
	bookController := &controllers.BookController{
		BookService: bookService,
	}

	router.POST("/books", bookController.CreateBook)
	router.GET("/books", bookController.ListBooks)
	router.GET("/books/:id", bookController.GetOneBook)
	router.PATCH("/books/:id", bookController.UpdateOneBook)
	router.DELETE("/books/:id", bookController.DeleteOneBook)
}
