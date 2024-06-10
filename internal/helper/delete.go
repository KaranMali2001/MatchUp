package helper

import (
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func DeleteInfo[T any](c echo.Context, db *gorm.DB, model *T, id string) error {

	resultChan := make(chan error)
	go func() {
		result := db.Where("id = ? ", id).Delete(model)
		if result.Error != nil {
			log.Println(result.Error)
			resultChan <- result.Error
			return
		}
		if result.RowsAffected > 0 {
			resultChan <- nil
			return
		} else {
			resultChan <- errors.New("record not found")
			return
		}

	}()
	err := <-resultChan
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, "deleted successfully")

}
