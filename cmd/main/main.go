package main

import (
	"github.com/KaranMali2001/MatchUp/internal/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORS())
	route.HomeRoute(e)
	route.Player_Route(e)
	route.Organizer_Route(e)
	route.Tournament_route(e)
	route.MatchRoute(e)
	e.Start(":8080")

}
