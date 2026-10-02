package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/sboy99/exp.go/go-auth/internal/models"
)

// -----------------------------INTERFACES-------------------------------- //

type IRoleRepository interface {
	Create(name string) (*models.Role, error)
	UpdatePermissions(roleId uint, permissions []string) (*models.Role, error)
}

// -----------------------------STRUCTS-------------------------------- //

type RoleRepository struct {
	DB *gorm.DB
}

// -----------------------------PUBLIC-------------------------------- //

func (r *RoleRepository) Create(name string) (*models.Role, error) {
	role := models.Role{
		Name:        name,
		Permissions: []string{},
	}

	err := r.DB.Create(&role).Error

	return &role, err
}

func (r *RoleRepository) UpdatePermissions(roleId uint, permissions []string) (*models.Role, error) {
	// todo: update permissions
	return nil, nil
}

// -----------------------------PRIVATE-------------------------------- //
