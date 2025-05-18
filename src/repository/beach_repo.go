package repository

import (
	"seaventures/src/config"
	"seaventures/src/models"
)

func CreateBeach(beach *models.Beach) error {
	return config.DB.Create(beach).Error
}

func GetAllBeaches() ([]models.Beach, error) {
	var beaches []models.Beach
	err := config.DB.Find(&beaches).Error
	return beaches, err
}

func GetBeachByID(id string) (*models.Beach, error) {
	var beach models.Beach
	err := config.DB.Where("id = ?", id).First(&beach).Error
	return &beach, err
}

func UpdateBeach(id string, updatedBeach *models.Beach) error {
	var beach models.Beach
	err := config.DB.Where("id = ?", id).First(&beach).Error
	if err != nil {
		return err
	}
	return config.DB.Model(&beach).Updates(updatedBeach).Error
}

func DeleteBeach(id string) error {
	return config.DB.Where("id = ?", id).Delete(&models.Beach{}).Error
}

func GetBeachesByLocationID(locationID string) ([]models.Beach, error) {
	var beaches []models.Beach
	err := config.DB.Where("l_id = ?", locationID).Find(&beaches).Error
	return beaches, err
}

func GetBeachDescriptionByBeachID(beachID string) (string, error) {
	var beach models.Beach
	err := config.DB.Select("beach_desc").Where("id = ?", beachID).First(&beach).Error
	if err != nil {
		return "", err
	}
	return beach.BeachDesc, nil
}


