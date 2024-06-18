package route

import (
	"github.com/KaranMali2001/MatchUp/internal/handlers/Match"
	"github.com/KaranMali2001/MatchUp/middleware"
	"github.com/labstack/echo"
)

func MatchRoute(e *echo.Echo) {
	e.GET("/match/:id", match.GetMatch)
	e.GET("/CreateMatch/:tournament_name", match.CreateMatch, middleware.VerifyJWT)
	e.PUT("/match/:id/score", match.UpdateScore, middleware.VerifyJWT)
	e.PUT("/match/:id/winner",match.UpdateWinner,middleware.VerifyJWT)
	e.GET("/match/:tournament_name", match.GetAllMatches)
}
