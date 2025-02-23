package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"userName"`
	Email    string `json:"email" gorm:"unique"`
	Role     string `json:"role"`
	ProfilePic string `json:"profilePic"`
}
