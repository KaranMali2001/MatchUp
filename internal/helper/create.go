package helper

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"log"
)

func Create[T any](db *gorm.DB, model *T) chan error {
	resultChan := make(chan error)
	fmt.Println("info about model",model)
	go func() {
		defer close(resultChan)
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
