 package tempdata

import (
	"fmt"
	"log"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
)
func Registration(){
	db:=database.Db
	var players []models.Player
	var Tournament models.Tournament
     var registration []models.Registration
	 if err:=db.Find(&registration).Error;err!=nil{
		log.Println(err)
		return
	 }
	 if(len(registration)>0){
		fmt.Println("registration are already added in db")
	return 
	}
	result:=db.Find(&players)
	if result.Error!=nil{
		log.Println(result.Error)
		return
	}
	if(len(players)<=0){
		fmt.Println("players dont exist")
		return 
	}
	result =db.Where("organizer_name = ?","DemoUser").First(&Tournament)
	if result.Error!=nil{
		log.Println(result.Error)
		return
	}
	for _,player:=range players{
		reg:=models.Registration{
			PlayerUsername: player.Username,
			TournamentName: Tournament.TournamentName,
		}
      if err:=db.Create(&reg).Error;err!=nil{
           log.Println(err)
		   return
	  }

	}
	fmt.Println("registration for tournament is successful")
	
} 