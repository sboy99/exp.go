package file

import (
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
)

// user repository structure
type UserFileRepository struct{}

// user repository methods
func (r *UserFileRepository) Create(name string, email string) *models.User {
	newUser := models.User{
		Name:  name,
		Email: email,
	}
	return &newUser
}
