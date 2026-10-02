package services

import (
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
)

type IAuthService interface {
	Register(name string, email string, password string) *models.User
	// Login(ctx *gin.Context) *models.User
}
