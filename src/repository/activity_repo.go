package repository

import (
	"seaventures/src/config"
	"seaventures/src/models"
)

func CreateActivity(activity *models.Activity) error {
	return config.DB.Create(activity).Error
}

func GetAllActivities() ([]models.Activity, error) {
	var activities []models.Activity
	err := config.DB.Find(&activities).Error
	return activities, err
}

func GetActivityByID(id string) (models.Activity, error) {
	var activity models.Activity
	err := config.DB.Where("id = ?", id).First(&activity).Error
	return activity, err
}

func UpdateActivity(id string, updatedActivity *models.Activity) error {
	var activity models.Activity
	err := config.DB.Where("id = ?", id).First(&activity).Error
	if err != nil {
		return err
	}
	return config.DB.Model(&activity).Updates(updatedActivity).Error
}

func DeleteActivity(id string) error {
	return config.DB.Where("id = ?", id).Delete(&models.Activity{}).Error
}

func GetActivityDescriptionByActivityID(id string) (string, error) {
	var activity models.Activity
	err := config.DB.Select("desc").Where("id = ?", id).First(&activity).Error
	if err != nil {
		return "", err
	}
	return activity.Desc, nil
}