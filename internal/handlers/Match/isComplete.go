package match

import (
	"fmt"
	"log"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
)

func isMatchCompleted(id int) string {
	var match models.Match
	db := database.Db
	
	err := db.Where("id = ?", id).Find(&match).Error
	if err != nil {
		fmt.Println("error occured",id)
		log.Println(err)
		return "error while finding match"
	}
	var incomplete int64
	tournament := match.TournamentName
	err = db.Model(&models.Match{}).Where("tournament_name = ? AND winner = ?", tournament, "").Count(&incomplete).Error
	if err != nil {
		log.Println(err)
		return "error while finding incomplete matches"
	}
	fmt.Println("inside isMatchesCompleted Fucntion ")
	if incomplete == 0 {
		return "success"
	}
	return "not completed"
}
