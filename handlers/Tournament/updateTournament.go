package tournament

import (
	"net/http"

	"github.com/labstack/echo"
)

func UpdateTournament(c echo.Context) error {
	return c.JSON(http.StatusOK, "updated tournament sucessfully")
}
