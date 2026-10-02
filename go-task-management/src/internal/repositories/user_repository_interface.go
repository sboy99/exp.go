package repositories

import "github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"

type IUserRepository interface {
	Create(slug string, name string, email string, password string) *models.User
	CreateSlug(name string) string
	GetUserByEmail(email string) *models.User
}
