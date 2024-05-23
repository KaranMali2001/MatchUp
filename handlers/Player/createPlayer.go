package player

import (
	"net/http"

	"github.com/labstack/echo"
)

func CreatePlayer(c echo.Context) error {
	return c.JSON(http.StatusOK, "player created sucessfully")
}
