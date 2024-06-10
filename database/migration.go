package database

import (
	"github.com/KaranMali2001/MatchUp/database/models"
)

func AutoMigrate() {
	Db.AutoMigrate(&models.Player{})
	Db.AutoMigrate(&models.Organizer{})
	Db.AutoMigrate(&models.Tournament{})
	Db.AutoMigrate(&models.Registration{})
	Db.AutoMigrate(&models.Match{})

}
func DropTable(){
	Db.Migrator().DropTable(&models.Match{})
}
