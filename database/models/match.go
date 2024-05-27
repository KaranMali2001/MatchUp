package models

import "gorm.io/gorm"

type Score struct {
	FirstPlayerScore  uint8 `json:"first_player_score"`
	SecondPlayerScore uint8 `json:"second_player_score"`
}
type Match struct {
	gorm.Model
	PlayerUsername1 string       `json:"player_username" gorm:"not null"`
	Player1         Registration `gorm:"foreignKey:PlayerUsername1;reference:PlayerUsername;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	PlayerUsername2 string       `json:"player_username2" gorm:"not null"`
	Player2         Registration `gorm:"foreignKey:PlayerUsername2;reference:PlayerUsername;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	TournamentName  string       `json:"tournament_name"`
	Tournament      Tournament   `gorm:"foreignKey:TournamentName;reference:TournamentName;constraint:onUpdate:CASCADE,onDelete:SET NULL"`
	SET1            Score        `json:"set_1" gorm:"embedded"`
	SET2            Score        `json:"set_2" gorm:"embedded"`
	SET3            Score        `json:"set_3" gorm:"embedded"`
}
