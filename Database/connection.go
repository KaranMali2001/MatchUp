package database

import (
	"fmt"

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
    Db.AutoMigrate(&Player{},&Organizer{})
	fmt.Println("database connected sucessfully")

}
