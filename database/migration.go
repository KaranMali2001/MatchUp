package database

import (
	"fmt"
	"log"

	"github.com/KaranMali2001/MatchUp/database/models"
)

func AutoMigrate() {
	err := Db.AutoMigrate(
        &models.Player{},
        &models.Organizer{},
        &models.Tournament{},
        &models.Registration{},
        &models.Match{},
    )
    if err != nil {
        log.Println("Error during migration:", err)
    } else {
        fmt.Println("Migration successful")
    }
	

}
