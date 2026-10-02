package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/sboy99/exp.go/go-task-management/src/internal/domain/models"
	"github.com/sboy99/exp.go/go-task-management/src/internal/helpers"
)

// user repository structure
type UserPostgresRepository struct {
	DB           *gorm.DB
	StringHelper helpers.IStrinigHelper
	CryptoHelper helpers.ICryptoHelper
}

func (r *UserPostgresRepository) Create(slug string, name string, email string, password string) *models.User {
	newUser := models.User{
		Slug:     slug,
		Name:     name,
		Email:    email,
		Password: password,
	}
	r.DB.Create(&newUser)
	return &newUser
}

func (r *UserPostgresRepository) CreateSlug(name string) string {
	slug := r.StringHelper.GetSlugString(name, nil)
	uniqSlug := slug

	retry := 10
	for i := 1; i <= retry; i++ {
		var user models.User
		err := r.DB.Where("slug = ?", uniqSlug).First(&user).Error
		if gorm.IsRecordNotFoundError(err) {
			return uniqSlug
		}
		uniqSlug = slug + r.CryptoHelper.GetRandStr(1000)
	}

	return ""
}

func (r *UserPostgresRepository) GetUserByEmail(email string) *models.User {
	var user models.User
	r.DB.Where("email = ?", email).First(&user)
	return &user
}
