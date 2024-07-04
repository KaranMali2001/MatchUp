package models

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	FirstName    string `json:"first_name" gorm:"not null"`
	LastName     string `json:"last_name" gorm:"not null"`
	Username     string `json:"username" gorm:"unique;not null" validate:"required"`
	Password     string `json:"password" validate:"required,min=8"`
	Email        string `json:"email" validate:"required,email"`
	Role         string `json:"role" gorm:"default:player"`
	TotalMatches int    `json:"total_matches"`
	Win          int    `json:"win"`
	Loss         int    `json:"loss"`
}
