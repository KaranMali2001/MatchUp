package match

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	fmt.Println("before tour live")
	if !tournament.Live {
		return c.JSON(http.StatusOK, "this tournament is not live")
	}
	if claims.Username != tournament.OrganizerName {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":    "this tournament is not created by you",
			"created by": tournament.OrganizerName,
		})
	}

	fmt.Println("before var players")
	var players []string
	for _, player := range reg {
		players = append(players, player.PlayerUsername)
	}

	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	fmt.Println(len(players))

	err = match(len(players), players)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while creating match")
	}

	return c.JSON(http.StatusOK, "matches created successfully")
}
func match(round int, players []string) error {
	if (len(players) % 2) != 0 {
		match := &models.Match{
			Round:                strconv.Itoa(round),
			FirstPlayerUsername:  players[0],
			SecondPlayerUsername: "",
		}

		err := database.Db.Create(match).Error
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("match for ", match.FirstPlayerUsername, match.SecondPlayerUsername, " created")
	}
	players = players[1:]
	for i := 0; i < len(players); i += 2 {
		match := &models.Match{
			Round:                strconv.Itoa(round),
			FirstPlayerUsername:  players[i],
			SecondPlayerUsername: players[i+1],
		}

		err := database.Db.Create(match).Error
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("match for ", match.FirstPlayerUsername, match.SecondPlayerUsername, " created")
	}

	return nil
}
