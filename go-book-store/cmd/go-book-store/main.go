package main

import (
	"github.com/sboy99/exp.go/go-book-store/adapters/repositories/file"
	"github.com/sboy99/exp.go/go-book-store/internal/routers"
)

func main() {
	// load data
	file.LoadBooks()

	// initialize all routes
	routers.Init()
}
