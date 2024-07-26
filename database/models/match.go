package models

import (
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	Round               string       `json:"round"`
	FirstPlayerUsername string       `json:"first_player_username" gorm:"type:varchar(100);not null"`
	PlayerOne           Registration `gorm:"foreignKey:FirstPlayerUsername;references:PlayerUsername;constraint:onUpdate:CASCADE,onDelete:SET NULL"`

	SecondPlayerUsername string `json:"second_player_username" gorm:"type:varchar(100)"`

	TournamentName string     `json:"tournament_name" gorm:"type:varchar(100);not null"`
	Tournament     Tournament `gorm:"foreignKey:TournamentName;references:TournamentName;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	Winner         string     `json:"winner"`
	PlayerOneSet1  uint       `json:"player_one_set1"`
	PlayerTwoSet1  uint       `json:"player_two_set1"`
	PlayerOneSet2  uint       `json:"player_one_set2"`
	PlayerTwoSet2  uint       `json:"player_two_set2"`
	PlayerOneSet3  uint       `json:"player_one_set3"`
	PlayerTwoSet3  uint       `json:"player_two_set3"`
}
