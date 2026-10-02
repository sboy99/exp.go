package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-todo/adapters/handlers/rest/transformers"
	"github.com/sboy99/exp.go/go-todo/infra/exception"
	"github.com/sboy99/exp.go/go-todo/internal/ports"
)

// ------------------------------INTERFACES---------------------------------- //

type IAuthRestHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

// ------------------------------STRUCTS---------------------------------- //

type AuthHandler struct {
	Service    ports.IAuthServicePort
	Tranformer transformers.IAuthTransformer
}

// ------------------------------METHODS---------------------------------- //

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	reqDto, err := h.Tranformer.GetRegisterReqDto(c)
	if err != nil {
		exception.HandleHttpException(c, err)
		return nil
	}

	user, err := h.Service.Register(reqDto.Email, reqDto.Pass, reqDto.Name)
	if err != nil {
		exception.HandleHttpException(c, err)
		return nil
	}

	resDto, err := h.Tranformer.GetRegisterResDto(user)
	if err != nil {
		exception.HandleHttpException(c, err)
		return nil
	}

	c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "User registered successfully",
		"user":    *resDto,
	})

	return nil
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	reqDto, err := h.Tranformer.GetLoginReqDto(c)
	if err != nil {
		exception.HandleHttpException(c, err)
		return nil
	}

	accessToken, refreshToken, err := h.Service.Login(reqDto.Email, reqDto.Pass)
	if err != nil {
		exception.HandleHttpException(c, err)
		return nil
	}

	// todo: transform

	c.Status(http.StatusOK).JSON(fiber.Map{
		"status":       http.StatusOK,
		"message":      "User registered successfully",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})

	return nil
}
