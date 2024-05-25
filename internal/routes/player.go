package route

import (
	middleware "github.com/KaranMali2001/MatchUp/middleware"

	player "github.com/KaranMali2001/MatchUp/internal/handlers/Player"
	"github.com/labstack/echo"
)

func Player_Route(e *echo.Echo) {
	e.GET("/player", player.GetAllPlayer)
	e.GET("/player/:username", player.GetPlayerInfo)
	e.POST("/player", player.CreatePlayer, middleware.Validator)
	e.PUT("/player/:username", player.UpdatePlayer, middleware.VerifyJWT)
	e.DELETE("/player/:username", player.DeletePlayer, middleware.VerifyJWT)
}
