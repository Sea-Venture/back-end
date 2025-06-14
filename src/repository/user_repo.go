package repository

import (
	"seaventures/src/config"
	"seaventures/src/models"
)

func Register(user *models.User) error {
	return config.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(userName string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(id uint) (*models.User, error) {
    var user models.User
    err := config.DB.Where("id = ?", id).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

func GetRoleByEmail(email string) (string, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

func UpdateUserRole(user *models.User, role string) error {
	if role != "admin" && role != "guide" && role != "user" {
		return config.DB.Error
	}
	user.Role = role
	return config.DB.Save(user).Error
}

func UpdateUserRoleById(id uint, role string) error {
	var user models.User
	err := config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	if role != "admin" && role != "guide" && role != "user" {
		return config.DB.Error
	}
	user.Role = role
	return config.DB.Save(&user).Error
}

func GetUserIdByEmail(email string) (uint, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}



