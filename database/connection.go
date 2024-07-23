package database

import (
	"fmt"
	"sync"

	"github.com/KaranMali2001/MatchUp/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	conn := "postgresql://postgres:password@localhost:5432/MatchUp"
	Db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
Db.AutoMigrate(models.Match{})
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
