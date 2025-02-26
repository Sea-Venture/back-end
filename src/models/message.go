package models

import (
	"time"
)

type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content"`
	Time      time.Time `json:"time"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
