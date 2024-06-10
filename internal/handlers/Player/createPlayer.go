package player

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/KaranMali2001/MatchUp/middleware"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func CreatePlayer(c echo.Context) error {
	db := database.Db
	player := new(models.Player)

	if err := c.Bind(&player); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while binding the body")
	}
	player.Role = "player"
	errChannel := helper.Create(db, player)
	// Wait for the result from the channel
	if err := <-errChannel; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "Error while creating tournament")
	}
	token, err := middleware.CreateToken("player", player.Username)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "error while creating jwt")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "player created successfully",
		"player":  player,
		"token":   token,
	})
}
