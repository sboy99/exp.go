package main

import (
	"log"

	"github.com/sboy99/exp.go/go-todo/app"
)

func main() {
	// create app
	app, err := app.Create()
	if err != nil {
		log.Fatal(err.Error())
	}

	// register middlwares
	app.RegisterMiddlewares()

	// register routes
	app.RegisterRoutes()

	// register database relations
	app.RegisterDatabaseRelations()

	// serve
	app.Serve()
}
