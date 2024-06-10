package route

import (
	"github.com/KaranMali2001/MatchUp/internal/handlers/Organizer"
	"github.com/KaranMali2001/MatchUp/middleware"
	"github.com/labstack/echo"
)

func Organizer_Route(e *echo.Echo) {
	e.GET("/org", organizer.GetAllOrg)
	e.GET("/organizer/:id", organizer.GetOrganizerInfo)
	e.POST("/organizer", organizer.NewOrganizer, middleware.Validator)
	e.PUT("/organizer/:id", organizer.UpdateOrganizer, middleware.VerifyJWT)
	e.DELETE("/organizer/:id", organizer.DeleteOrganizer, middleware.VerifyJWT)
}
