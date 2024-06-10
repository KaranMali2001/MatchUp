package models

import (
	"gorm.io/gorm"
)
type Match struct {
    gorm.Model

    FirstPlayerUsername  string       `json:"first_player_username" gorm:"type:varchar(100);not null"`
    PlayerOne            Registration `gorm:"foreignKey:FirstPlayerUsername;references:PlayerUsername;constraint:onUpdate:CASCADE,onDelete:SET NULL"`

    SecondPlayerUsername string       `json:"second_player_username" gorm:"type:varchar(100);not null"`
    PlayerTwo            Registration `gorm:"foreignKey:SecondPlayerUsername;references:PlayerUsername;constraint:onUpdate:CASCADE,onDelete:SET NULL"`

    TournamentName       string       `json:"tournament_name" gorm:"type:varchar(100);not null"`
    Tournament           Tournament   `gorm:"foreignKey:TournamentName;references:TournamentName;constraint:onUpdate:CASCADE,onDelete:SET NULL"`

    SET1                 Score        `json:"set_1" gorm:"embedded"`
    SET2                 Score        `json:"set_2" gorm:"embedded"`
    SET3                 Score        `json:"set_3" gorm:"embedded"`
}

type Score struct {
    FirstPlayerScore  uint8 `json:"first_player_score"`
    SecondPlayerScore uint8 `json:"second_player_score"`
}
