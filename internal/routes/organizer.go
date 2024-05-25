package route

import (
	middleware "github.com/KaranMali2001/MatchUp/Middleware"
	organizer "github.com/KaranMali2001/MatchUp/internal/handlers/Organizer"
	"github.com/labstack/echo"
)

func Organizer_Route(e *echo.Echo) {
	e.GET("/org", organizer.GetAllOrg)
	e.GET("/organizer/:username", organizer.GetOrganizerInfo)
	e.POST("/organizer", organizer.NewOrganizer, middleware.Validator)
	e.PUT("/organizer/:username", organizer.UpdateOrganizer)
	e.DELETE("/organizer/:username", organizer.DeleteOranizer)
}
