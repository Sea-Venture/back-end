package service

import (
	"seaventures/src/models"
	"seaventures/src/repository"

)

func CreateBeach(beach *models.Beach) error {
	return repository.CreateBeach(beach)
}

func GetAllBeaches() ([]models.Beach, error) {
	return repository.GetAllBeaches()
}

func GetBeachByID(id string) (*models.Beach, error) {
	return repository.GetBeachByID(id)
}

func UpdateBeach(id string, updatedBeach *models.Beach) error {
	return repository.UpdateBeach(id, updatedBeach)
}

func DeleteBeach(id string) error {
	return repository.DeleteBeach(id)
}

func GetBeachesByLocationID(locationID string) ([]models.Beach, error) {
	return repository.GetBeachesByLocationID(locationID)
}

func GetBeachDescriptionByBeachID(beachID string) (string, error) {
	return repository.GetBeachDescriptionByBeachID(beachID)
}