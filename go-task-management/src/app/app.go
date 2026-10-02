package app

import (
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/controllers/transformers"
	"github.com/sboy99/exp.go/go-task-management/src/adapters/repositories/postgres"
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/exp.go/go-task-management/src/internal/helpers"
	"github.com/sboy99/exp.go/go-task-management/src/internal/services"
	"github.com/sboy99/exp.go/go-task-management/src/routers"
)

func (app *App) RegisterRouters() {
	// helpers
	cryptoHelper := &helpers.CryptoHelper{}
	stringHelper := &helpers.StringHelper{}

	// repositories
	userPostgresRepo := &postgres.UserPostgresRepository{
		DB:           app.DB,
		CryptoHelper: cryptoHelper,
		StringHelper: stringHelper,
	}
	taskPostgresRepo := &postgres.TaskPostgresRepository{
		DB: app.DB,
	}

	// services
	authService := &services.AuthService{
		UserRepo:     userPostgresRepo,
		CryptoHelper: cryptoHelper,
	}
	userService := &services.UserService{
		UserRepo:     userPostgresRepo,
		CryptoHelper: cryptoHelper,
	}
	taskService := &services.TaskService{
		TaskRepo: taskPostgresRepo,
	}

	// transformers
	userTransformer := &transformers.UserTransformer{}

	// controllers
	authController := &controllers.AuthController{
		Service:         authService,
		UserTransformer: userTransformer,
	}
	userController := &controllers.UserController{
		Service:     userService,
		Transformer: userTransformer,
	}
	taskController := &controllers.TaskController{
		Service: taskService,
	}

	// routers
	authRouter := &routers.AuthRouter{
		Router:     app.Router,
		Controller: authController,
	}
	userRouter := &routers.UserRouter{
		Router:     app.Router,
		Controller: userController,
	}
	taskRouter := &routers.TaskRouter{
		Router:     app.Router,
		Controller: taskController,
	}

	// register routes
	authRouter.Register()
	userRouter.Register()
	taskRouter.Register()
}

func (app *App) RegisterRelations() {
	app.DB.AutoMigrate(&models.User{}, &models.Task{})
}

func (app *App) Start() {
	app.Router.Run("localhost:3000")
}
