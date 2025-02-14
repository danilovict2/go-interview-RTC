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

func (r *UserRepository) FindByUUID(uuid string) (models.User, error) {
    user := models.User{}
	if err := r.DB.First(&user, "uuid = ?", uuid).Error; err != nil {
        return models.User{}, err
    }

    return user, nil
}