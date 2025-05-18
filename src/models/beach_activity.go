package models

import "gorm.io/gorm"

type BeachActivity struct {
	gorm.Model
	BeachID    uint `json:"beach_id"`
	ActivityID uint `json:"activity_id"`
}
