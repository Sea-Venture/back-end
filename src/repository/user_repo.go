package repository

import (
	"seaventures/src/config"
	"seaventures/src/models"
)

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}
