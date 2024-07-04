package main

import (
	"fmt"

	"github.com/KaranMali2001/MatchUp/internal/routes"
	//tempdata "github.com/KaranMali2001/MatchUp/tempData"

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
	fmt.Println("dummy data is not loading")
	route.DummyData()

	e.Start(":8080")

}
