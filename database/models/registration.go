package models

import (
	"gorm.io/gorm"
)

type Registration struct {
	gorm.Model
	PlayerUsername string     `json:"player_username" gorm:"not null"`
	Player         Player     `gorm:"foreignKey:PlayerUsername;references:Username;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	TournamentName string     `json:"tournament_name" gorm:"not null"`
	Tournament     Tournament `gorm:"foreignKey:TournamentName;references:TournamentName;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
}
