package services

import (
	"seaventures/src/models"
	"seaventures/src/repository"
)

func CreateUserService(user *models.User) error {
	return repository.CreateUser(user)
}

func GetUsersService() ([]models.User, error) {
	return repository.GetUsers()
}

func UpdateUserService(user *models.User) error {
	return repository.UpdateUser(user)
}