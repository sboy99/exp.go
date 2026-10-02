package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers/transformers"
	"github.com/sboy99/exp.go/go-task-management/src/internal/services"
)

// user controller structure
type UserController struct {
	Transformer transformers.IUserTransformer
	Service     services.IUserService
}

// user controller methods
func (c *UserController) Create(ctx *gin.Context) {
	dto := c.Transformer.TransformCreateReq(ctx)
	userData := c.Service.Create(dto.Name, dto.Email, dto.Password)
	user := c.Transformer.TransformCreateRes(userData)
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) List(ctx *gin.Context) {
	// todo: transform request body
}

func (c *UserController) GetById(ctx *gin.Context) {
	// todo: transform request body
}

func (c *UserController) UpdateById(ctx *gin.Context) {
	// todo: transform request body
}

func (c *UserController) DeleteById(ctx *gin.Context) {
	// todo: transform request body
}
