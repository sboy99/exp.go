package repositories

import "github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"

type ITaskRepository interface {
	Create(creatorId uint, name string, priority int) *models.Task
	// List() *[]models.Task
	// GetById(id uint) *models.Task
	// updateById(id uint) *models.Task
	// DeleteById(id uint) *models.Task
}
