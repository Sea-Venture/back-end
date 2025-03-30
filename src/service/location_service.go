package service

import (
	"seaventures/src/models"
	"seaventures/src/repository"

)




func CreateLocation(location *models.Location) error {
	return repository.CreateLocation(location)
}

func  GetLocationByID(id string) (*models.Location, error) {
    return repository.GetLocationByID(id)
}

func GetLocations() ([]models.Location, error) {
	return repository.GetLocations()
}

func UpdateLocation(id string, location *models.Location) error {
	return repository.UpdateLocation(id, location)
}

func DeleteLocation(id string) error {
	return repository.DeleteLocation(id)
}




