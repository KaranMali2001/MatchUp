package route

import (
	organizer "github.com/KaranMali2001/MatchUp/internal/handlers/Organizer"
	middleware "github.com/KaranMali2001/MatchUp/middleware"
	"github.com/labstack/echo"
)

func Organizer_Route(e *echo.Echo) {
	e.GET("/org", organizer.GetAllOrg)
	e.GET("/organizer/:username", organizer.GetOrganizerInfo)
	e.POST("/organizer", organizer.NewOrganizer, middleware.Validator)
	e.PUT("/organizer/:username", organizer.UpdateOrganizer, middleware.VerifyJWT)
	e.DELETE("/organizer/:username", organizer.DeleteOrganizer, middleware.VerifyJWT)
}
