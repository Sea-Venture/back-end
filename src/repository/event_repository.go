package repository

import (
	"seaventures/src/config"
	"seaventures/src/models"
)

func CreateEvent(event *models.Event) error {
	return config.DB.Create(event).Error
}

func GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	err := config.DB.Find(&events).Error
	return events, err
}

func GetEventByID(id string) (models.Event, error) {
	var event models.Event
	err := config.DB.Where("id = ?", id).First(&event).Error
	return event, err
}

func UpdateEvent(id string, updatedEvent *models.Event) error {
	var event models.Event
	err := config.DB.Where("id = ?", id).First(&event).Error
	if err != nil {
		return err
	}
	return config.DB.Model(&event).Updates(updatedEvent).Error
}

func DeleteEvent(id string) error {
	return config.DB.Where("id = ?", id).Delete(&models.Event{}).Error
}

func GetEventByActivityID(activityID string) ([]models.Event, error) {
	var events []models.Event
	err := config.DB.Where("activity_id = ?", activityID).Find(&events).Error
	return events, err
}

func GetEventByLocationID(locationID string) ([]models.Event, error) {
	var events []models.Event
	err := config.DB.Where("location_id = ?", locationID).Find(&events).Error
	return events, err
}

func GetEvenByLocationIDAndActivityID(locationID string, activityID string) ([]models.Event, error) {
	var events []models.Event
	err := config.DB.Where("location_id = ? AND activity_id = ?", locationID, activityID).Find(&events).Error
	return events, err
}
