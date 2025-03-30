package service

import (
	"seaventures/src/models"
	"seaventures/src/repository"

)

func CreateActivity(activity *models.Activity) error {
	return repository.CreateActivity(activity)
}

func GetActivityByID(id string) (models.Activity, error) {
	return repository.GetActivityByID(id)
}


func GetAllActivities() ([]models.Activity, error) {
	return repository.GetAllActivities()
}

func UpdateActivity(id string, activity *models.Activity) error {
	return repository.UpdateActivity(id, activity)
}

func DeleteActivity(id string) error {
	return repository.DeleteActivity(id)
}

func GetActivityDescriptionByActivityID(id string) (string, error) {
	return repository.GetActivityDescriptionByActivityID(id)
}