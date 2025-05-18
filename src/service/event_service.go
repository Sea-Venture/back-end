package service

import (
	"seaventures/src/models"
	"seaventures/src/repository"
)

func CreateEvent(event *models.Event) error {
	return repository.CreateEvent(event)
}

func GetAllEvents() ([]models.Event, error) {
	return repository.GetAllEvents()
}


func GetEventByID(id string) (models.Event, error) {
	return repository.GetEventByID(id)
}

func UpdateEvent(id string, blog *models.Event) error {
	return repository.UpdateEvent(id, blog)
}

func DeleteEvent(id string) error {
	return repository.DeleteEvent(id)
}

func GetEventByActivityID(activityID string) ([]models.Event, error) {
	return repository.GetEventByActivityID(activityID)
}

func GetEventByLocationID(locationID string) ([]models.Event, error) {
	return repository.GetEventByLocationID(locationID)
}

func GetEventByLocationIDAndActivityID(locationID string, activityID string) ([]models.Event, error) {
	return repository.GetEvenByLocationIDAndActivityID(locationID, activityID)
}
