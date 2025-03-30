package models

import (
	"time"
)

type Activity struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
