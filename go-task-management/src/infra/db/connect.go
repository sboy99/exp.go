package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sboy99/exp.go/go-task-management/src/infra/exceptions"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=godb sslmode=disable password=postgres")
	// handle database connection exceptions
	exceptions.HandleDBConnectionException(err)
	// done
	return db
}
