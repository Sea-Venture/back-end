package repository

import (
	"seaventures/src/config"
	"seaventures/src/models"
)

func CreateGuide(guide *models.Guide) error {
	return config.DB.Create(guide).Error
}

func GetAllGuides() ([]models.Guide, error) {
	var guides []models.Guide
	err := config.DB.Find(&guides).Error
	return guides, err
}

func GetGuideByID(id string) (*models.Guide, error) {
	var guide models.Guide
	err := config.DB.Where("id = ?", id).First(&guide).Error
	return &guide, err
}

func UpdateGuide(id string, updatedGuide *models.Guide) error {
	var guide models.Guide
	err := config.DB.Where("id = ?", id).First(&guide).Error
	if err != nil {
		return err
	}
	return config.DB.Model(&guide).Updates(updatedGuide).Error
}

func DeleteGuide(id string) error {
	return config.DB.Where("id = ?", id).Delete(&models.Guide{}).Error
}

func GetGuideByBeachID(beachID string) ([]models.Guide, error) {
	var guides []models.Guide
	err := config.DB.Where("beach_id = ?", beachID).Find(&guides).Error
	return guides, err
}

func GetGuideByActivityID(activityID string) ([]models.Guide, error) {
    var guides []models.Guide

    err := config.DB.Joins("JOIN guide_activities ON guide_activities.guide_id = guides.id").
        Where("guide_activities.activity_id = ?", activityID).
        Find(&guides).Error

    return guides, err
}

func GetGuideActivitiesByBeachIDAndActivityID(beachID string, activityID string) ([]models.Guide, error) {
	var guides []models.Guide

	err := config.DB.Joins("JOIN guide_activities ON guide_activities.guide_id = guides.id").
		Where("guides.beach_id = ? AND guide_activities.activity_id = ?", beachID, activityID).
		Find(&guides).Error

	return guides, err
}