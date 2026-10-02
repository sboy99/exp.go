package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-todo/internal/ports"
)

// ------------------------------INTERFACES---------------------------------- //

type IHealthRestHandler interface {
	Check(c *fiber.Ctx) error
}

// ------------------------------STRUCTS---------------------------------- //

type HealthHandler struct {
	Service ports.IHealthServicePort
}

// ------------------------------METHODS---------------------------------- //

func (h *HealthHandler) Check(c *fiber.Ctx) error {
	msg, err := h.Service.Check()
	c.Status(http.StatusOK).JSON(fiber.Map{
		"message": *msg,
	})
	return err
}
