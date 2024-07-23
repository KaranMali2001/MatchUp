package main

import (
	"fmt"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/internal/routes"

	//tempdata "github.com/KaranMali2001/MatchUp/tempData"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"}, // Change this to your Next.js frontend URL
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},

		AllowCredentials: true,
	}))
	route.HomeRoute(e)
	route.Player_Route(e)
	route.Organizer_Route(e)
	route.Tournament_route(e)
	route.MatchRoute(e)
	fmt.Println("dummy data is  loading")
	fmt.Println("added cookie to code ")
	route.DummyData()
	e.GET("/DELETE-DATA", database.DeleteData)
	e.Start(":8080")
	//fix match update issue
	//only first set is updating
	//docker run -e POSTGRES_PASSWORD=password -e POSTGRES_DB=MatchUp -v MatchUpDB:/var/lib/postgresql/data -p 5432:5432 postgres

}
