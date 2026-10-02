package services

import (
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/exp.go/go-task-management/src/internal/helpers"
	"github.com/sboy99/exp.go/go-task-management/src/internal/repositories"
)

// user service structures
type UserService struct {
	UserRepo     repositories.IUserRepository
	CryptoHelper helpers.ICryptoHelper
}

// user service methods
func (s *UserService) Create(name string, email string, password string) *models.User {
	userSlug := s.UserRepo.CreateSlug(name)
	hashedPassword := s.CryptoHelper.GetHash(password)
	return s.UserRepo.Create(userSlug, name, email, hashedPassword)
}

func (s *UserService) List() *[]models.User {
	return nil
}

func (s *UserService) GetById(id int64) *models.User {
	return nil
}

func (s *UserService) UpdateById(id int64) *models.User {
	return nil
}

func (s *UserService) DeleteById(id int64) *models.User {
	return nil
}
