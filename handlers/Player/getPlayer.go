package player

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetPlayerInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, "returning player info")
}
