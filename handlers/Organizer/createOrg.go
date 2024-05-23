package organizer

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewOrganizer(c echo.Context) error {
	return c.JSON(http.StatusOK, "new org created sucessfully")
}
