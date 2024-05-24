package route

import (
	tournament "github.com/KaranMali2001/MatchUp/handlers/Tournament"
	"github.com/labstack/echo"
)

func Tournament_route(e *echo.Echo) {
	e.GET("/Tournament", tournament.AllTournament)
	e.POST("/Tournament", tournament.NewTournament)
	e.PUT("/Tournament/:tournament_name", tournament.UpdateTournament)
	e.DELETE("/Tournament/:tournament_name", tournament.DeleteTournament)
	e.GET("/Tournament/:tournament_name", tournament.GetTournament)
}
