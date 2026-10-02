package ports

import (
	"github.com/sboy99/exp.go/go-todo/internal/domain/models"
)

// -------------------------------INTERFACES--------------------------------- //

type ITaskRepositoryPort interface {
	Create(name string, priority int) (*models.Task, error)
}

type ITaskServicePort interface {
	Create(name string, priority int) (*models.Task, error)
}

// type TaskRepository struct {
// 	DB *gorm.DB
// }
