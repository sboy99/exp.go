package services

import (
	"github.com/sboy99/exp.go/go-task-management/src/internal/cache"
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/exp.go/go-task-management/src/internal/helpers"
	"github.com/sboy99/exp.go/go-task-management/src/internal/repositories"
)

type AuthService struct {
	UserRepo     repositories.IUserRepository
	SessionCache cache.ISessionCache
	CryptoHelper helpers.ICryptoHelper
}

func (s *AuthService) Register(name string, email string, password string) *models.User {
	// check for existing emails
	s._checkForRegisteredEmail(email)
	// create user
	user := s._createUser(name, email, password)

	// create active session

	// return user
	return user
}

func (s *AuthService) _checkForRegisteredEmail(email string) {
	user := s.UserRepo.GetUserByEmail(email)

	if user != nil {
		// todo: throw http exception
		return
	}
}

func (s *AuthService) _createUser(name string, email string, password string) *models.User {
	// get user slug
	userSlug := s.UserRepo.CreateSlug(name)
	// get hash password
	hashedPassword := s.CryptoHelper.GetHash(password)
	// create a new user
	return s.UserRepo.Create(userSlug, name, email, hashedPassword)
}

func (s *AuthService) _createActiveSession(user *models.User) {
	// todo: create active session with id
}
