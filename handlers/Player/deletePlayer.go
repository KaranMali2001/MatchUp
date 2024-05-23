package player

import (
	"net/http"

	"github.com/labstack/echo"
)

func DeletePlayer(c echo.Context) error {
	return c.JSON(http.StatusOK, "player deleted sucessfullly")
}
