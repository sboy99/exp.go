package services

import (
	"github.com/sboy99/exp.go/go-todo/internal/ports"
)

// ------------------------------STRUCTS---------------------------------- //

type UserService struct {
	UserRepo ports.IUserRepositoryPort
}

// ------------------------------PUBLIC_METHODS---------------------------------- //
