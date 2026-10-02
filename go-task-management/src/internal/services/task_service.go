package services

import (
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/exp.go/go-task-management/src/internal/repositories"
)

type TaskService struct {
	TaskRepo repositories.ITaskRepository
}

func (s *TaskService) Create(creatorId uint, name string, priority int) *models.Task {
	return s.TaskRepo.Create(creatorId, name, priority)
}
