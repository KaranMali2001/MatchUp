package homepage

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "set-cookie"
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.Path = "/"
	cookie.MaxAge = -1
	fmt.Println("cookie in broswer are", cookie)
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "log out successful")
}
