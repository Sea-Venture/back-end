package service

import (
	"seaventures/src/models"
	"seaventures/src/repository"

)

func CreateGuide(guide *models.Guide) error {
	return repository.CreateGuide(guide)
}

func GetAllGuides() ([]models.Guide, error) {
	return repository.GetAllGuides()
}

func GetGuideByID(id string) (*models.Guide, error) {
	return repository.GetGuideByID(id)
}

func UpdateGuide(id string, updatedGuide *models.Guide) error {
	return repository.UpdateGuide(id, updatedGuide)
}

func DeleteGuide(id string) error {
	return repository.DeleteGuide(id)
}

func GetGuideByBeachID(beachID string) ([]models.Guide, error) {
	return repository.GetGuideByBeachID(beachID)
}

func GetGuideByActivityID(activityID string) ([]models.Guide, error) {
	return repository.GetGuideByActivityID(activityID)
}

func GetGuideActivitiesByBeachIDAndActivityID(beachID string, activityID string) ([]models.Guide, error) {
	return repository.GetGuideActivitiesByBeachIDAndActivityID(beachID, activityID)
}