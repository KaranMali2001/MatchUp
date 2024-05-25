package main

import (
	"fmt"

	database "github.com/KaranMali2001/MatchUp/database"
	route "github.com/KaranMali2001/MatchUp/internal/routes"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	route.Player_Route(e)
	route.Organizer_Route(e)
	route.Tournament_route(e)

	e.Start(":8080")

	fmt.Println(database.Db)
}
