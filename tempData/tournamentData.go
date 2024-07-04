package tempdata
import (
	"fmt"
	"log"
	

	

"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"

	
	
)


func  TournamentData()  {
	db:=database.Db
	var Tournament []models.Tournament
	err:=db.Find(&Tournament).Error
	if err != nil {
		log.Println(err)
	}
	if(len(Tournament)>0){
	    fmt.Println("data exist already for tournament")
		return
	}
	temp:=models.Tournament{
		TournamentName: "demo tournament",
		TotalPlayer: 0,
		OrganizerName: "DemoUser",
		Live: true,
	}
	
	
   
	
	if err:=db.Create(&temp).Error; err != nil {
		
		log.Println(err)
		return 
	

}
fmt.Println("tournament created successfully")
	
}
