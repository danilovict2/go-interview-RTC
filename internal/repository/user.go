package repository

import (
	"github.com/danilovict2/go-interview-RTC/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (r *UserRepository) FindOneByUUID(uuid string) (models.User, error) {
	user := models.User{}
	if err := r.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) FindByUUID(uuids []string) ([]models.User, error) {
	users := make([]models.User, 0)
	for _, uuid := range uuids {
		user, err := r.FindOneByUUID(uuid)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}
