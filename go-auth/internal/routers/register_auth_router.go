package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-auth/internal/transport/rest"
)

func RegisterAuthRouter(app *fiber.App) {
	app.Post("/auth/register", rest.HandleRegister)
	app.Post("/auth/login", rest.HandleLogin)
}
