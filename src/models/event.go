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
	LocationID  uint       `json:"location_id"`
	ActivityID  uint       `json:"activity_id"`
}
