package services

import (
	"github.com/sboy99/exp.go/go-todo/internal/domain/models"
	"github.com/sboy99/exp.go/go-todo/internal/ports"
)

// ------------------------------STRUCTS---------------------------------- //

type TaskService struct {
	TaskRepo ports.ITaskRepositoryPort
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (s *TaskService) Create(name string, priority int) (*models.Task, error) {
	return s.TaskRepo.Create(name, priority)
}
