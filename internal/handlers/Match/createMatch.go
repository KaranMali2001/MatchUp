package match

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
)

func CreateMatch(c echo.Context) error {
	tournament_name := c.Param("tournament_name")
	db := database.Db
	claims := c.Get("claims").(*models.JWTClaims)
	
	var reg []models.Registration
	fmt.Println("tournament name is " + tournament_name)
	err := db.Where("tournament_name = ?", tournament_name).Find(&reg).Error
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while finding this tournament")
	}
	
	var tournament models.Tournament
	err = db.Where("tournament_name= ?", tournament_name).First(&tournament).Error
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while finding tournament")
	}
	
	if claims.Username != tournament.OrganizerName {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":    "this tournament is not created by you",
			"created by": tournament.OrganizerName,
		})
	}
	var playerName []string
	for _, names := range reg {
		playerName = append(playerName, names.PlayerUsername)
	fmt.Println(names.PlayerUsername)
	}
	fmt.Println("name of all player ....")
	 res, err := shuffleMatches(playerName, tournament_name)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while shuffling ")
	} 
	return c.JSON(http.StatusOK, res)
}

func shuffleMatches(playerName []string, tournamentName string) ([][]string, error) {
	db := database.Db

	rand.Shuffle(len(playerName), func(i, j int) {
		playerName[i], playerName[j] = playerName[j], playerName[i]
	})
	var pairs [][]string
	for i := 0; i < len(playerName); i += 2 {
		pair := []string{playerName[i], playerName[i+1]}
		pairs = append(pairs, pair)
	}
	for _, match := range pairs {
		newMatch := &models.Match{
			
			TournamentName:  tournamentName,
			PlayerUsername1: match[0],
			PlayerUsername2: match[1],
		}
		if err := db.Create(&newMatch).Error; err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return pairs, nil
}
