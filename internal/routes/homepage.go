package route

import (
	"net/http"

	homepage "github.com/KaranMali2001/MatchUp/internal/handlers/HomePage"
	"github.com/labstack/echo"
)
func HomeRoute(e *echo.Echo){
	e.GET("/",greating)
	
	e.POST("/login",homepage.Login)
}
func greating(c echo.Context)error{
	return c.JSON(http.StatusOK,"server started")
}