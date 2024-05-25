package models
import (
	"gorm.io/gorm"
)
type Registration struct {
	gorm.Model
	PlayerUsername string     `json:"player_username" gorm:"not null;index"`
	Player         Player     `gorm:"foreignKey:PlayerUsername;references:Username;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	TournamentID   string     `json:"tournament_id" gorm:"not null"`
	Tournament     Tournament `gorm:"foreignKey:TournamentID;references:TournamentID;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
}