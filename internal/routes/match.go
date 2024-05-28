package route

import (
	match "github.com/KaranMali2001/MatchUp/internal/handlers/Match"
	"github.com/KaranMali2001/MatchUp/middleware"
	"github.com/labstack/echo"
)

func MatchRoute(e *echo.Echo) {
	e.GET("/match/:id", match.GetMatch)
	e.POST("/match/:tournament_name", match.CreateMatch, middleware.VerifyJWT)
	e.PUT("/match/:id", match.UpdateMatch, middleware.VerifyJWT)
	e.GET("/match/:tournament_name", match.GetAllMatches)
}
