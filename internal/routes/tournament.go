package route

import (
	tournament "github.com/KaranMali2001/MatchUp/internal/handlers/Tournament"
	"github.com/KaranMali2001/MatchUp/middleware"
	"github.com/labstack/echo"
)

func Tournament_route(e *echo.Echo) {
	e.GET("/Tournament", tournament.AllTournament)
	e.POST("/Tournament", tournament.NewTournament, middleware.VerifyJWT)
	e.PUT("/Tournament/:tournament_name", tournament.UpdateTournament, middleware.VerifyJWT)
	e.DELETE("/Tournament/:tournament_name", tournament.DeleteTournament, middleware.VerifyJWT)
	e.GET("/Tournament/:tournament_name", tournament.GetTournament)
}