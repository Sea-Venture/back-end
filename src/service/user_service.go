package service

import (
	"errors"
	"seaventures/src/models"
	"seaventures/src/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(user *models.User) error {

	existingUser, err := repository.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email is already taken")
	}

	existingUser, err = repository.GetUserByUsername(user.UserName)
	if err == nil && existingUser != nil {
		return errors.New("username is already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Role = "user"

	return repository.Register(user)
}

func Login(user *models.User) error {
	existingUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return errors.New("invalid credentials")
	}

	user.ID = existingUser.ID
	user.Role = existingUser.Role

	return nil
}

func UpdateProfilePic(user *models.User, profilePic string) error {

	user.ProfilePic = profilePic

	return repository.UpdateUser(user)
}

func GetUserById(id uint) (*models.User, error) {
	return repository.GetUserById(id)
}
