package route

import (
	middleware "github.com/KaranMali2001/MatchUp/middleware"

	"github.com/KaranMali2001/MatchUp/internal/handlers/Player"
	"github.com/labstack/echo"
)

func Player_Route(e *echo.Echo) {
	e.GET("/player", player.GetAllPlayer)
	//need to fix this
	e.GET("/player/:id", player.GetPlayerInfo)
	e.POST("/player", player.CreatePlayer, middleware.Validator)
	//need to fix both of them
	e.PUT("/player/:id", player.UpdatePlayer, middleware.VerifyJWT)
	e.DELETE("/player/:id", player.DeletePlayer, middleware.VerifyJWT)
}
