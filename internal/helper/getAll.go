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
			resultChan <- gorm.ErrRecordNotFound
			return
		}
		resultChan <- nil
	}()

	err := <-resultChan
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "record not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

	}
	return c.JSON(http.StatusOK,models)
}

