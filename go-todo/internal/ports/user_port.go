package ports

import (
	"github.com/sboy99/exp.go/go-todo/internal/domain/models"
)

// -------------------------------INTERFACES--------------------------------- //

type IUserRepositoryPort interface {
	Create(email string, pass string, name string, slug string) (*models.User, error)
	GetSlug(name string) (string, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUser(conds *models.User) (*models.User, error)
}

type IUserServicePort interface{}
