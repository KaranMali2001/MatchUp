package tournament

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewTournament(c echo.Context) error {
	return c.JSON(http.StatusOK, "new tournament created sucessfully")
}
