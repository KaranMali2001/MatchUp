package models

import (
	"gorm.io/gorm"
	"time"
)

type Tournament struct {
	gorm.Model

	TournamentName string    `json:"tournament_name" gorm:"not null;unique"`
	StartDate      time.Time `json:"start_date" gorm:"not null"`
	EndDate        time.Time `json:"end_date" gorm:"not null"`
	TotalPlayer    int       `json:"total_player"`
	OrganizerName  string    `json:"organizer_name" gorm:"not null"`
	Organizer      Organizer `gorm:"foreignKey:OrganizerName;references:Username;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Live           bool      `json:"live" gorm:"default:false"`
	WinnerOfTournament string `json:"winner_of_tournament"`
}
