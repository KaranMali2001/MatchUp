package tournament

import (
	"net/http"

	"github.com/labstack/echo"
)

func AllTournament(c echo.Context) error {
	return c.JSON(http.StatusOK, "returning all the tournaments")
}
