package match

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/labstack/echo"
	"gorm.io/gorm"
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

	for _, player := range reg {
		playerName = append(playerName, player.PlayerUsername)

		fmt.Println(player.PlayerUsername)
	}
	playerCount := len(playerName)
	fmt.Println(playerCount)
	rounds := (playerCount - 1) / 2
	tx := db.Begin()
	if playerCount%2 == 1 {

		randIndex := rand.Intn(playerCount)
		newMatch := &models.Match{
			Round:                fmt.Sprintf("round of %d", rounds),
			FirstPlayerUsername:  playerName[randIndex],
			SecondPlayerUsername: "bye",
			TournamentName:       tournament_name,
			SET1:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
			SET2:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
			SET3:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
			Winner:               playerName[randIndex],
		}
		errChan := make(chan error, 1)
		var l1 sync.Mutex
		go func() {

			defer func() {
				if err := recover(); err != nil {

					l1.Lock()
					errChan <- fmt.Errorf("recovered from error")
					l1.Unlock()
				}
			}()
			if err := tx.Create(&newMatch).Error; err != nil {

				l1.Lock()
				errChan <- err
				l1.Unlock()

			}

		}()
		err := <-errChan
		if err != nil {
			tx.Rollback() // Rollback if creation fails
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, "error while creating matches")
		}
		playerName = append(playerName[:randIndex], playerName[randIndex+1:]...)
	}

	fmt.Println("name of all player ....")
	res, err := shuffleMatches(tx, playerName, tournament_name)
	if err != nil {
		tx.Rollback()
		log.Println(err)

	}
	tx.Commit()
	return c.JSON(http.StatusOK, res)
}

func shuffleMatches(tx *gorm.DB, playerName []string, tournamentName string) ([][]string, error) {

	rand.Shuffle(len(playerName), func(i, j int) {
		playerName[i], playerName[j] = playerName[j], playerName[i]
	})
	var pairs [][]string
	for i := 0; i < len(playerName); i += 2 {
		pair := []string{playerName[i], playerName[i+1]}
		pairs = append(pairs, pair)
	}
	fmt.Println("pairs", pairs)
	rounds := (len(playerName) - 1) / 2
	for _, pair := range pairs {
		newMatch := &models.Match{
			Round: fmt.Sprintf("round of %d", rounds),

			FirstPlayerUsername:  pair[0],
			SecondPlayerUsername: pair[1],
			TournamentName:       tournamentName,
			SET1:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
			SET2:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
			SET3:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
		}
		var errChan chan error
		var lock sync.Mutex
		go func() {
			if err := tx.Create(&newMatch).Error; err != nil {

				lock.Lock()
				errChan <- err
				lock.Unlock()
			}

		}()
		err := <-errChan
		if err != nil {
			return nil, err
		}
	}
	return pairs, nil
}
