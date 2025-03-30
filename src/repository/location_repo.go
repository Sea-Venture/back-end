package repository

import (
	"seaventures/src/models"
	"seaventures/src/config"
)

func CreateLocation(location *models.Location) error {
	return config.DB.Create(location).Error
}

func GetLocationByID(id string) (*models.Location, error) {
	var location models.Location
	err := config.DB.Where("id = ?", id).First(&location).Error
	return &location, err
}

func GetLocations() ([]models.Location, error) {
	var locations []models.Location
	err := config.DB.Find(&locations).Error
	return locations, err
}

func UpdateLocation(id string, updatedLocation *models.Location) error {
	var location models.Location
	err := config.DB.Where("id = ?", id).First(&location).Error
	if err != nil {
		return err
	}
	return config.DB.Model(&location).Updates(updatedLocation).Error
}

func DeleteLocation(id string) error {
	return config.DB.Where("id = ?", id).Delete(&models.Location{}).Error
}