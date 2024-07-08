package tournament

import (
	"log"
	"net/http"
	"sync"
"fmt"
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func Registration(c echo.Context) error {
	reg := new(models.Registration)
	db := database.Db
	var tournament models.Tournament

	TournamentName := c.Param("tournament_name")
	claims := c.Get("claims").(*models.JWTClaims)
	if claims.Role != "player" {
		return c.JSON(http.StatusUnauthorized, "you are not player")
	}
	if err := db.Model(&tournament).Where("live = ?", true).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "tournament is not live")
	}
	reg.TournamentName = TournamentName
	reg.PlayerUsername = claims.Username
	errChannel := helper.Create(db, reg)
	if err := <-errChannel; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while creating")
	}
	var mutex sync.Mutex
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("error while incrementing player")
			}
		}()
		mutex.Lock()
		defer mutex.Unlock()
		tournament.TotalPlayer += tournament.TotalPlayer + 1
		if err := db.Save(&tournament); err != nil {
			log.Println(err)

		}
		
		
	}()
	go func() {
		var player models.Player
		err := db.Where("username =?", reg.PlayerUsername).Find(&player).Error
		if err != nil {
			log.Println(err)
		    fmt.Println("error while finding player with given username")			 
			return
		}
         err=db.Model(&player).UpdateColumn("total_matches",player.TotalMatches+1).Error
		 if err != nil {
			log.Println(err)
			fmt.Println("error while updating total matches in player")
			return
		 }
		}()

	
	return c.JSON(http.StatusOK, "registration for tournament is successful")
}
