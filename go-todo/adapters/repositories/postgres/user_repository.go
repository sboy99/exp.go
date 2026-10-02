package postgres

import (
	"errors"

	"github.com/sboy99/exp.go/go-todo/internal/domain/models"
	"github.com/sboy99/exp.go/go-todo/internal/helpers"
	"gorm.io/gorm"
)

// ------------------------------STRUCTS---------------------------------- //

type UserPostgresRepository struct {
	DB *gorm.DB
}

// ------------------------------PUBLIC_METHODS---------------------------------- //

func (r *UserPostgresRepository) Create(email string, pass string, name string, slug string) (*models.User, error) {
	// create a new task
	user := &models.User{
		Slug:  slug,
		Name:  name,
		Email: email,
		Pass:  pass,
	}
	// save the task in database
	err := r.DB.Create(user).Error
	return user, err
}

func (r *UserPostgresRepository) GetSlug(name string) (string, error) {
	user := new(models.User)
	delimeter := ""
	slug := helpers.GenSlugStr(name, delimeter)

	for true {
		err := r.DB.First(user, &models.User{
			Slug: slug,
		}).Error
		if err != nil {
			return slug, nil
		}

		slug += delimeter + helpers.GetSlugSalt(1000)

	}

	return "", errors.New("Can't generate slug")
}

func (r *UserPostgresRepository) GetUserByEmail(email string) (*models.User, error) {
	user, err := r.GetUser(&models.User{Email: email})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserPostgresRepository) GetUser(conds *models.User) (*models.User, error) {
	user := new(models.User)

	err := r.DB.First(user, conds).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// ------------------------------PRIVATE_METHODS---------------------------------- //
