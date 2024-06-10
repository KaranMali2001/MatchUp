package helper

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetInfo[T any](c echo.Context, db *gorm.DB, model *T, id uint64) error {
	resultChan := make(chan error)
	fmt.Println(id)
	go func() {
		result := db.Where("id = ? ", id).First(&model)
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
	return c.JSON(http.StatusOK, model)
}
