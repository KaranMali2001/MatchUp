package helper

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func UpdateInfo[T any](c echo.Context, db *gorm.DB, model *T, id int) error {
	err := db.First(model, "id = ?", id).Error
	resultChan := make(chan error)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while finding ")
	}
	if err := c.Bind(model); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	go func() {
		result := db.Save(model)
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

	error1 := <-resultChan
	if error1 != nil {
		log.Println(error1)
		return c.JSON(http.StatusInternalServerError, "error")
	}

	return c.JSON(http.StatusOK, "record has been updated")
}
