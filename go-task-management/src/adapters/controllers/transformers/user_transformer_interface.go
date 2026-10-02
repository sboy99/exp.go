package transformers

import (
	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers/dtos"
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
)

type IUserTransformer interface {
	TransformCreateReq(ctx *gin.Context) dtos.CreateUserReqDto
	TransformCreateRes(user *models.User) dtos.CreateUserResDto
}
