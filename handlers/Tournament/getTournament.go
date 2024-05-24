package tournament

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetTournament(c echo.Context) error {
	return c.JSON(http.StatusOK, "returning specific tournament")
}
