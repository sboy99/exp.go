package cache

import "github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"

type ISessionCache interface {
	Get(key []string) *models.Session
	Set(key []string, value *models.Session)
}
