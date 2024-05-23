package organizer

import (
	"net/http"

	"github.com/labstack/echo"
)

func UpdateOrganizer(c echo.Context) error {
	return c.JSON(http.StatusOK, "org info updated sucessfully")
}
