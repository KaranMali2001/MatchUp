package tournament

import (
	"fmt"
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"sync"
)

func Registration(c echo.Context) error {
	reg := new(models.Registration)
	db := database.Db
	

	TournamentName := c.Param("tournament_name")
	claims := c.Get("claims").(*models.JWTClaims)
	if claims.Role != "player" {
		return c.JSON(http.StatusUnauthorized, "you are not player")
	}
	if err := db.Where("live = ? AND tournament_name = ?",true,TournamentName).Error; err != nil {
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
				return
			}
		}()
		mutex.Lock()
		defer mutex.Unlock()
		var tournament models.Tournament
		if err:= db.Where("tournament_name = ?",TournamentName).Find(&tournament).Error;err!=nil{
			log.Println(err)
			return 
		}
		fmt.Println(tournament)
		tournament.TotalPlayer += 1
		if err := db.Save(tournament); err != nil {
			log.Println(err)
			return

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
		err = db.Model(&player).UpdateColumn("total_matches", player.TotalMatches+1).Error
		if err != nil {
			log.Println(err)
			fmt.Println("error while updating total matches in player")
			return
		}
	}()

	return c.JSON(http.StatusOK, "registration for tournament is successful")
}
