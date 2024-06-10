package helper

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetAllInfo[T any](c echo.Context, db *gorm.DB, models *[]T) error {
	resultChan := make(chan error)
	go func() {
		result := db.Find(models)
		if result.Error != nil {
			log.Println(result.Error)
			resultChan <- result.Error
			return
		}
		if result.RowsAffected == 0 {
			resultChan <- errors.New("record not found")
			return
		}
		resultChan <- nil
	}()
	err := <-resultChan
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, models)
}
