package services

import "github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"

type IUserService interface {
	Create(name string, email string, password string) *models.User
	List() *[]models.User
	GetById(id int64) *models.User
	UpdateById(id int64) *models.User
	DeleteById(id int64) *models.User
}
