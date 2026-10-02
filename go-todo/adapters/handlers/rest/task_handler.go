package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-todo/adapters/handlers/rest/dtos"
	"github.com/sboy99/exp.go/go-todo/internal/ports"
)

// ------------------------------INTERFACES---------------------------------- //

type ITaskRestHandler interface {
	Create(c *fiber.Ctx) error
}

// ------------------------------STRUCTS---------------------------------- //

type TaskHandler struct {
	Service ports.ITaskRepositoryPort
}

// ------------------------------METHODS---------------------------------- //

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	createTaskDto := new(dtos.CreateTaskDto)
	err := c.BodyParser(createTaskDto)
	if err != nil {
		panic(err.Error())
	}

	task, err := h.Service.Create(createTaskDto.Name, createTaskDto.Priority)
	if err != nil {
		panic(err.Error())
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"data": *task,
	})
	return err
}
