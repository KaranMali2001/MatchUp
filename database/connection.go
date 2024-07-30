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
	//conn := "postgresql://postgres:password@postgres:5432/MatchUp"

	//	conn := "postgresql://postgres:password@localhost:5432/MatchUp"
	conn:="postgres://avnadmin:AVNS_3Hp-HoMwD06UlWXBfv-@pg-1ed982b1-karansignup5599.f.aivencloud.com:13094/defaultdb?sslmode=require"
	
	//docker run -e POSTGRES_PASSWORD=password -e POSTGRES_DB=MatchUp -v MatchUpDB:/var/lib/postgresql/data -p 5432:5432 -d postgres

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
