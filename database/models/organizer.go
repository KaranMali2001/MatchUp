package models

import (
	"gorm.io/gorm"
)

type Organizer struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username" gorm:"unique;not null" validate:"required"`
	Password  string `json:"password" validate:"required,min=8"`
	Email     string `json:"email" validate:"required,email"`
	Role      string `json:"role" gorm:"default:organizer"`
}
