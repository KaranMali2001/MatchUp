package match

import (
	"fmt"
	"log"
 "math/rand"
	"github.com/KaranMali2001/MatchUp/database/models"
	"gorm.io/gorm"
)
//TOD
//handle odd player for second round
//merge createMatch n this function same code repeated
func Rounds(db *gorm.DB,id string){
	var incompleteMatches int64
	var tournament_name string
	var match models.Match
	err:=db.Where("id = ?",id).First(&match).Error
	if err != nil {
		log.Println("error while finding tournament",err)
return
	}
	tournament_name=match.TournamentName
	round:=match.Round
    err = db.Model(&models.Match{}).Where("winner = ? AND tournament_name=?", "",tournament_name).Count(&incompleteMatches).Error
	if err != nil {
		log.Println(err)
		return
	}
	if incompleteMatches!=0{
		fmt.Println("matches are not completed yet")
		return
	}
	var matches []models.Match
	err=db.Where("tournament_name=? AND round = ?",tournament_name,round).Find(&matches).Error
	if err != nil {
		log.Println(err)
		return
	}
	var playerName []string
	for _,v:=range matches{
		playerName=append(playerName, v.Winner)
	}
	rand.Shuffle(len(playerName), func(i, j int) {
		playerName[i], playerName[j] = playerName[j], playerName[i]
	})
	var pairs [][]string
	for i := 0; i < len(playerName); i += 2 {
		pair := []string{playerName[i], playerName[i+1]}
		pairs = append(pairs, pair)
	}
	fmt.Println("pairs", pairs)
	rounds := (len(playerName) - 1) / 2
	for _, pair := range pairs {
		newMatch := &models.Match{
			Round: fmt.Sprintf("round of %d", rounds),

			FirstPlayerUsername:  pair[0],
			SecondPlayerUsername: pair[1],
			TournamentName:       tournament_name,
			SET1:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
			SET2:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
			SET3:                 models.Score{FirstPlayerScore: 0, SecondPlayerScore: 0},
		}
	
		go func() {
			if err := db.Create(&newMatch).Error; err != nil {

			log.Println(err)
			return
			}

		}()
		
	}
}