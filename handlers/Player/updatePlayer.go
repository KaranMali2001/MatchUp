package player

import (
	"net/http"

	"github.com/labstack/echo"
)

func UpdatePlayer(c echo.Context) error {
	return c.JSON(http.StatusOK, "player updated sucessfully")
}
