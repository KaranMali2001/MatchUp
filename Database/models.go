package database

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" gorm:"default:player"`
}
type Organizer struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" gorm:"default:organizer"`
}
type Tournament struct {
	gorm.Model
	TournamentID string `json:"tournament_id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	TournamentName string    `json:"tournament_name" gorm:"not null"`
	StartDate      time.Time `json:"start_date" gorm:"not null"`
	EndDate        time.Time `json:"end_date" gorm:"not null"`
	TotalPlayer    int       `json:"total_player"`
	OrganizerName  string    `json:"organizer_name" gorm:"not null;index"`
	Organizer      Organizer `gorm:"foreignKey:OrganizerName;references:Username;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Live           bool      `json:"live" gorm:"default:false"`
}
