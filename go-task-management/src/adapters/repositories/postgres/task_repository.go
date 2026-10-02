package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
)

// user repository structure
type TaskPostgresRepository struct {
	DB *gorm.DB
}

func (r *TaskPostgresRepository) Create(creatorId uint, name string, priority int) *models.Task {
	newTask := models.Task{
		Name:     name,
		Priority: priority,
		IsDone:   false,
		UserId:   creatorId,
	}
	r.DB.Create(&newTask)
	return &newTask
}
