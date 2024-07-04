package player

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/KaranMali2001/MatchUp/middleware"

	"github.com/labstack/echo"
)

func CreatePlayer(c echo.Context) error {
	db := database.Db
	player := &models.Player{
		Role:         "player",
		TotalMatches: 0,
		Win:          0,
		Loss:         0,
	}
	if err := c.Bind(&player); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while binding the body")
	}
	errChannel := helper.Create(db, player)
	// Wait for the result from the channel
	err := <-errChannel
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error while creating player")
	}
	str, err := middleware.CreateToken(c, player.Role, player.Username)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(str)
	return c.JSON(http.StatusOK, "Your account as player has been created and cookie has been set")
}
