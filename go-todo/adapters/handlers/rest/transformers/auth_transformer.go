package transformers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sboy99/exp.go/go-todo/adapters/handlers/rest/dtos"
	"github.com/sboy99/exp.go/go-todo/infra/exception"
	"github.com/sboy99/exp.go/go-todo/internal/domain/models"
)

// -------------------------------INTERFACE--------------------------------- //

type IAuthTransformer interface {
	GetRegisterReqDto(c *fiber.Ctx) (*dtos.RegisterReqDto, *exception.HttpException)
	GetRegisterResDto(user *models.User) (*dtos.RegisterResDto, *exception.HttpException)
	GetLoginReqDto(c *fiber.Ctx) (*dtos.LoginReqDto, *exception.HttpException)
	// GetLoginResDto(user *models.User) (*dtos.LoginResDto, *exception.HttpException)
}

// -------------------------------STRUCT--------------------------------- //

type AuthTransformer struct{}

// -------------------------------PUBLIC_METHODS--------------------------------- //

func (t *AuthTransformer) GetRegisterReqDto(c *fiber.Ctx) (*dtos.RegisterReqDto, *exception.HttpException) {
	dto := new(dtos.RegisterReqDto)
	err := c.BodyParser(&dto)
	if err != nil {
		return nil, exception.BadRequestException(err)
	}

	return dto, nil
}

func (t *AuthTransformer) GetRegisterResDto(user *models.User) (*dtos.RegisterResDto, *exception.HttpException) {
	dto := &dtos.RegisterResDto{
		Slug: user.Slug,
		Name: user.Name,
	}

	return dto, nil
}

func (t *AuthTransformer) GetLoginReqDto(c *fiber.Ctx) (*dtos.LoginReqDto, *exception.HttpException) {
	dto := new(dtos.LoginReqDto)
	err := c.BodyParser(&dto)
	if err != nil {
		return nil, exception.BadRequestException(err)
	}

	return dto, nil
}
