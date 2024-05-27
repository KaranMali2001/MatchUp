package database

import (
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/KaranMali2001/MatchUp/database/models"
)

func UpdateStatus() {
	cron := cron.New()
	cron.AddFunc("@daily", func() {
		db := Db
		var tournaments []models.Tournament
		if err := db.Find(&tournaments); err != nil {
			log.Println(err)
			return
		}
		currentTime := time.Now()
		for _, tournament := range tournaments {
			if currentTime.After(tournament.EndDate) {
				tournament.Live = false
				if err := db.Save(tournament); err != nil {
					log.Println(err)
					return
				}
			}
		}
	})

	cron.Start()
}
