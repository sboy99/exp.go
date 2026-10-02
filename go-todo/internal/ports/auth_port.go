package ports

import (
	"github.com/sboy99/exp.go/go-todo/infra/exception"
	"github.com/sboy99/exp.go/go-todo/internal/domain/models"
)

// -------------------------------SERVICE--------------------------------- //

type IAuthStrategyPort interface {
	Authenticate(strategy IJwtExtractorStrategyPort, payload interface{}) error
	ExtractToken(strategy IJwtExtractorStrategyPort) (string, error)
}

// -------------------------------SERVICE--------------------------------- //

type IAuthServicePort interface {
	Register(email string, pass string, name string) (*models.User, *exception.HttpException)
	Login(email string, pass string) (string, string, *exception.HttpException)
}
