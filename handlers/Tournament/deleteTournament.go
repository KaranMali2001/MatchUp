package tournament

import (
	"net/http"

	"github.com/labstack/echo"
)

func DeleteTournament(c echo.Context) error {
	return c.JSON(http.StatusOK, "deleted tournament sucessfully")
}
