package match

import (
	//"fmt"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"

	//"github.com/KaranMali2001/MatchUp/internal/helper"
	"github.com/labstack/echo"
)

func UpdateScore(c echo.Context) error {
    id := c.Param("id")
    idint, err := strconv.Atoi(id)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid match ID"})
    }

    var newScore struct {
        Set_1_first_player int `json:"set_1_first_player"`

        Set_1_second_player int `json:"set_1_second_player"`

        Set_2_first_player int `json:"set_2_first_player"`

         Set_2_second_player int`json:"set_2_second_player"`
        Set_3_first_player int`json:"set_3_first_player"`
 Set_3_second_player int `json:"set_3_second_player"`
    }

    if err := c.Bind(&newScore); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
    }

    var match models.Match
    db := database.Db

    if err := db.First(&match, idint).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Match not found"})
    }

    match.SET1.FirstPlayerScore = uint8(newScore.Set_1_first_player)
    match.SET1.SecondPlayerScore = uint8(newScore.Set_1_second_player)
    match.SET2.FirstPlayerScore=uint8(newScore.Set_2_first_player)
    fmt.Println( match.SET2.FirstPlayerScore)
    match.SET2.SecondPlayerScore=uint8(newScore.Set_2_second_player)
    match.SET3.FirstPlayerScore=uint8(newScore.Set_3_first_player)
    match.SET3.SecondPlayerScore=uint8(newScore.Set_3_second_player)

    if err := db.Updates(&match).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not update match"})
    }

    return c.JSON(http.StatusOK, match)
}
