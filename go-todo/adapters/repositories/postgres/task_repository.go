package postgres

import (
	"github.com/sboy99/exp.go/go-todo/internal/domain/models"
	"gorm.io/gorm"
)

// ------------------------------STRUCTS---------------------------------- //

type TaskPostgresRepository struct {
	DB *gorm.DB
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (r *TaskPostgresRepository) Create(name string, priority int) (*models.Task, error) {
	// create a new task
	task := &models.Task{
		Name:     name,
		Priority: priority,
		UserId:   1,
	}
	// save the task in database
	err := r.DB.Create(task).Error
	return task, err
}
