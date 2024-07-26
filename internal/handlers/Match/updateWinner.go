package match

import (
	//"fmt"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
	"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateWinner(c echo.Context) error {
	idstr := c.Param("id")
	id, _:= strconv.Atoi(idstr)
	var match models.Match

	db := database.Db
err:= helper.UpdateInfo(c, db, &match, id)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, "error while updating")

	}
	
	var round int
	round,_= strconv.Atoi(match.Round)
	round=round/2
	tournamentName:=match.TournamentName
	errChan2:=make(chan error)

	
	go func() {
		
		defer close(errChan2)
		str := isMatchCompleted(id)
	
		if str == "success" {
			fmt.Println("all matches are completed for tournament ",tournamentName," & matches for round no ",round," will begin")
			
				var players []string
				err:= db.Model(&models.Match{}).Where("tournament_name = ? AND winner IS NOT NULL AND round = ?",tournamentName,match.Round).Pluck("winner",&players).Error
				if err != nil {
					log.Println(err)
					errChan2 <- err
                    return
				}
				fmt.Println("winner of round :",match.Round,"are ",players)

			  err=Match(round,players,tournamentName)	
			  if err != nil {
				log.Println(err)

				errChan2 <- err
				return 
			  }
			  errChan2<-nil 
			
		} else {
			fmt.Println(str)
		   errChan2<-nil
		}
	}()
	err=<-errChan2
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,"match has been updated but not able to prepare next round")
	} 
	return c.JSON(http.StatusOK, "match has been updated successfully")
}
