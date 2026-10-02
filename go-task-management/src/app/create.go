package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
	"github.com/sboy99/exp.go/go-task-management/src/infra/cache"
	"github.com/sboy99/exp.go/go-task-management/src/infra/db"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	Cache  *redis.Client
}

func Create() *App {
	// get router
	router := gin.Default()

	// connect to database
	db := db.Connect()

	// connect to cache
	cache := cache.Connect()

	return &App{
		Router: router,
		DB:     db,
		Cache:  cache,
	}
}
