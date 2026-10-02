package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers/transformers"
	"github.com/sboy99/exp.go/go-task-management/src/internal/services"
)

type AuthController struct {
	Service         services.IAuthService
	UserTransformer transformers.IUserTransformer
}

func (c *AuthController) Register(ctx *gin.Context) {
	dto := c.UserTransformer.TransformCreateReq(ctx)
	user := c.Service.Register(dto.Name, dto.Email, dto.Password)
	userRes := c.UserTransformer.TransformCreateRes(user)
	ctx.JSON(http.StatusOK, userRes)
}
