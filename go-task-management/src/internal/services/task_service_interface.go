package services

import "github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"

type ITaskService interface {
	Create(creatorId uint, name string, priority int) *models.Task
}
