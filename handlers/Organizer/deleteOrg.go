package organizer

import (
	"net/http"

	"github.com/labstack/echo"
)

func DeleteOranizer(c echo.Context) error {
	return c.JSON(http.StatusOK, "org deleted sucessfully")
}
