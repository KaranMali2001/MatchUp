package database

import (
	"fmt"
	"sync"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	conn := "postgresql://postgres:password@postgres:5432/MatchUp"
	Db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	AutoMigrate()
	

	fmt.Println("database connected successfully")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		UpdateStatus()
		wg.Done()
	}()
	wg.Wait()
}
