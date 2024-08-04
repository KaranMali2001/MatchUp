package main

import (
	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/internal/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},

		AllowCredentials: true,
	}))
	route.HomeRoute(e)
	route.Player_Route(e)
	route.Organizer_Route(e)
	route.Tournament_route(e)
	route.MatchRoute(e)

	route.DummyData()
	e.GET("/DELETE-DATA", database.DeleteData)
	e.Start(":8080")

}
