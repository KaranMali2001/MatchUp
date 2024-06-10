package helper

import (
	"errors"

	"gorm.io/gorm"

	"log"
)

func Create[T any](db *gorm.DB, model *T) chan error {
	resultChan := make(chan error)
	go func() {
		result := db.Create(model)
		if result.Error != nil {
			log.Println(result.Error)
			resultChan <- result.Error
			return
		}
		if result.RowsAffected > 0 {
			resultChan <- nil
			return
		} else {
			resultChan <- errors.New("failed to create db record")
			return
		}

	}()
	return resultChan

}
