package route

import (
	"github.com/KaranMali2001/MatchUp/internal/handlers/Tournament"
	"github.com/KaranMali2001/MatchUp/middleware"
	"github.com/labstack/echo"
)

func Tournament_route(e *echo.Echo) {
	e.GET("/Tournament", tournament.AllTournament)
	e.POST("/Tournament", tournament.NewTournament, middleware.VerifyJWT)
	e.PUT("/Tournament/:id", tournament.UpdateTournament, middleware.VerifyJWT)
	e.DELETE("/Tournament/:id", tournament.DeleteTournament, middleware.VerifyJWT)
	e.GET("/Tournament/:id", tournament.GetTournament)
	e.GET("/Tournament/registration/:tournament_name", tournament.Registration, middleware.VerifyJWT)
}
