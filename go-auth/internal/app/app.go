package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-auth/internal/routers"
)

var app *fiber.App

// ------------------------------PRIVATE---------------------------------- //

func create() {
	app = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "Go Auth v1.0.0",
	})
}

func initRoutes() {
	routers.RegisterAuthRouter(app)
}

func initMiddlewares() {
	// todo: implement
}

func serve() {
	app.Listen(":8080")
}

// ------------------------------PUBLIC---------------------------------- //

func Bootstrap() {
	create()
	initRoutes()
	initMiddlewares()
	serve()
}
