package tempdata

import (
	"fmt"
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"log"
	"sync"
)

func Registration() {
	db := database.Db
	var players []models.Player
	var Tournament models.Tournament
	var registration []models.Registration
	if err := db.Find(&registration).Error; err != nil {
		log.Println(err)
		return
	}
	if len(registration) > 0 {
		fmt.Println("registration are already added in db")
		return
	}
	result := db.Find(&players)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	if len(players) <= 0 {
		fmt.Println("players dont exist")
		return
	}
	result = db.Where("organizer_name = ?", "DemoUser").First(&Tournament)
	if result.Error != nil {
		log.Println(result.Error)
		return
	}
	for _, player := range players {
		reg := models.Registration{
			PlayerUsername: player.Username,
			TournamentName: Tournament.TournamentName,
		}
		if err := db.Create(&reg).Error; err != nil {
			log.Println(err)
			return
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("error while incrementing player")
				}
			}()
			var mutex sync.Mutex
			mutex.Lock()
			defer mutex.Unlock()
			Tournament.TotalPlayer += +1

			if err := db.Save(&Tournament).Error; err != nil {
				log.Println(err)
				fmt.Println("error while updating total player")
				return

			}

		}()
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("recovered from panic:", r)
				}
			}()
			var player models.Player
			err := db.Where("username =?", reg.PlayerUsername).Find(&player).Error
			if err != nil {
				log.Println(err)
				fmt.Println("error while finding player with given username")
				return
			}
			err = db.Model(&player).UpdateColumn("total_matches", player.TotalMatches+1).Error
			if err != nil {
				log.Println(err)
				fmt.Println("error while updating total matches in player")
				return
			}
		}()

	}
	fmt.Println("registration for tournament is successful")

}
