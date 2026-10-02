package app

import (
	"github.com/sboy99/exp.go/go-todo/adapters/handlers/rest"
	"github.com/sboy99/exp.go/go-todo/adapters/handlers/rest/transformers"
	"github.com/sboy99/exp.go/go-todo/adapters/jwt"
	"github.com/sboy99/exp.go/go-todo/adapters/repositories/postgres"
	"github.com/sboy99/exp.go/go-todo/internal/helpers"
	"github.com/sboy99/exp.go/go-todo/internal/services"
	"gorm.io/gorm"
)

// ------------------------------STRUCTS---------------------------------- //

type Handler struct {
	DB *gorm.DB
}

type RestHandlers struct {
	AuthHandler   rest.IAuthRestHandler
	HealthHandler rest.IHealthRestHandler
	TaskHandler   rest.ITaskRestHandler
}

// ------------------------------PRIVATE---------------------------------- //

func (h *Handler) initRestHandlers() *RestHandlers {
	// repositories
	userRepo := &postgres.UserPostgresRepository{
		DB: h.DB,
	}
	taskRepo := &postgres.TaskPostgresRepository{
		DB: h.DB,
	}

	// helpers
	cryptoHelper := &helpers.CryptoHelper{}

	// adapters
	jwtAdapter := &jwt.JwtAdapter{}

	// services
	healthService := &services.HealthService{}
	authService := &services.AuthService{
		UserRepo:     userRepo,
		JwtAdapter:   jwtAdapter,
		CryptoHelper: cryptoHelper,
	}
	taskService := &services.TaskService{
		TaskRepo: taskRepo,
	}

	// transformers
	authTransformer := &transformers.AuthTransformer{}

	// handlers
	healthHandler := &rest.HealthHandler{
		Service: healthService,
	}
	authHandler := &rest.AuthHandler{
		Service:    authService,
		Tranformer: authTransformer,
	}
	taskHandler := &rest.TaskHandler{
		Service: taskService,
	}

	// done
	return &RestHandlers{
		HealthHandler: healthHandler,
		AuthHandler:   authHandler,
		TaskHandler:   taskHandler,
	}
}
