package homepage

import (
	"net/http"

	"github.com/labstack/echo"
)

func LogOut(c echo.Context) error {
	cookie := http.Cookie{
		Name:  "set-cookie",
		Value: "",
	}
	c.SetCookie(&cookie)
	return c.JSON(http.StatusOK, "log out successful")
}
