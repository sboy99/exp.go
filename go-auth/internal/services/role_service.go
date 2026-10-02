package services

import (
	"github.com/sboy99/exp.go/go-auth/internal/models"
	"github.com/sboy99/exp.go/go-auth/internal/repositories"
)

// -----------------------------INTERFACES-------------------------------- //

type IRoleService interface {
	Create(name string) (*models.Role, error)
	UpdatePermissions(roleId uint, permissions []string) (*models.Role, error)
}

// -----------------------------STRUCTS-------------------------------- //

type RoleService struct {
	RoleRepo repositories.IRoleRepository
}

// -----------------------------PUBLIC-------------------------------- //

func (s *RoleService) Create(name string) (*models.Role, error) {
	return s.RoleRepo.Create(name)
}

func (s *RoleService) UpdatePermissions(roleId uint, permissions []string) (*models.Role, error) {
	return s.RoleRepo.UpdatePermissions(roleId, permissions)
}
