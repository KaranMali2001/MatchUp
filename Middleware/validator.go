package middleware

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func Validator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		v := validator.New()
		err := v.Struct(c.Request().Body)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, "struct is not validated")
		}
		return next(c)
	}
}
