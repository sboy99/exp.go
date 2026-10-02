package transformers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers/dtos"
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
)

type UserTransformer struct{}

func (t *UserTransformer) TransformCreateReq(ctx *gin.Context) dtos.CreateUserReqDto {
	var createUserReqDto dtos.CreateUserReqDto
	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.Decode(&createUserReqDto)
	return createUserReqDto
}

func (t *UserTransformer) TransformCreateRes(user *models.User) dtos.CreateUserResDto {
	createUserResDto := dtos.CreateUserResDto{
		Id:        user.ID,
		Slug:      user.Slug,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return createUserResDto
}
