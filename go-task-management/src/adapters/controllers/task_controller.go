package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sboy99/exp.go/go-task-management/src/internal/services"
)

type TaskController struct {
	Service services.ITaskService
}

func (c *TaskController) Create(ctx *gin.Context) {
	task := c.Service.Create(1, "Read Book", 2)
	ctx.JSON(http.StatusOK, *task)
}
