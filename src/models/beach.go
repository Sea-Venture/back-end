package models

import "gorm.io/gorm"

type Beach struct {
	gorm.Model
	BeachID    string     `json:"beach_id"`
	BeachName  string     `json:"beach_name"`
	BeachDesc  string     `json:"beach_desc"`
	BeachType  string     `json:"beach_type"`
	LID        uint       `json:"location_id"`
	Activities []Activity `gorm:"many2many:beach_activities" json:"activities"` // Many-to-Many relationship
}
