package models

import "gorm.io/gorm"

type GuideActivity struct {
	gorm.Model
	GuideID    uint `json:"guide_id"`
	ActivityID uint `json:"activity_id"`
}
