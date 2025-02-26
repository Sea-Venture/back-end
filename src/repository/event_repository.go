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