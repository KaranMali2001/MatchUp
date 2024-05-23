package organizer

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetOrganizerInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, "this is organization info")
}
