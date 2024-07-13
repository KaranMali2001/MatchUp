package route

import (
	"net/http"

	homepage "github.com/KaranMali2001/MatchUp/internal/handlers/HomePage"
	"github.com/labstack/echo"
)

func HomeRoute(e *echo.Echo) {
	e.GET("/", greeting)

	e.POST("/login", homepage.Login)
	e.GET("/logout", homepage.LogOut)
}
func greeting(c echo.Context) error {
	return c.JSON(http.StatusOK, "server started")
}
