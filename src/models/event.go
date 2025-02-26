package models

import (
	"time"
)

type Event struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	SafeDelete  bool       `json:"safe_delete"`
	Activities  []Activity `gorm:"foreignKey:EventID" json:"activities"`
	LocationID  uint       `json:"location_id"`
}
