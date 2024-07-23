package tempdata

import (
	"fmt"
	"log"

	"github.com/KaranMali2001/MatchUp/database"
	"github.com/KaranMali2001/MatchUp/database/models"
)

func PlayerData() {
	db := database.Db
	var players []models.Player
	err := db.Find(&players)
	if err != nil {
		log.Println(err)
	}
	if len(players) > 0 {
		fmt.Println("data exist already for player")
		return
	}
	Players := []models.Player{
		
		{FirstName: "Karan", LastName: "Sharma", Username: "karan", Password: "password123", Email: "karan@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Arushi", LastName: "Mehta", Username: "arushi", Password: "password123", Email: "arushi@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Sakshi", LastName: "Verma", Username: "sakshi", Password: "password123", Email: "sakshi@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Ansh", LastName: "Patel", Username: "ansh", Password: "password123", Email: "ansh@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Program", LastName: "Singh", Username: "program", Password: "password123", Email: "program@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Yashavi", LastName: "Rao", Username: "yashavi", Password: "password123", Email: "yashavi@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Kashvi", LastName: "Gupta", Username: "kashvi", Password: "password123", Email: "kashvi@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Shruti", LastName: "Malhotra", Username: "shruti", Password: "password123", Email: "shruti@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Anchal", LastName: "Kumar", Username: "anchal", Password: "password123", Email: "anchal@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Lidia", LastName: "Fernandes", Username: "lidia", Password: "password123", Email: "lidia@example.com", TotalMatches: 0, Win: 0, Loss: 0},
		{FirstName: "Payal", LastName: "Desai", Username: "payal", Password: "password123", Email: "payal@example.com", TotalMatches: 0, Win: 0, Loss: 0},
	}
	if err := db.CreateInBatches(Players, 20); err != nil {
		log.Println(err)

	}
	fmt.Println("player data added successfully")
}

func OrganizerData() {
	db := database.Db
	var organizer []models.Organizer
	err := db.Find(&organizer)
	if err != nil {
		log.Println(err)
	}
	if len(organizer) > 0 {
		fmt.Println("data exist already for org")
		return
	}
	organizers := []models.Organizer{
		{FirstName: "demo", LastName: "user", Username: "DemoUser", Password: "demo123", Email: "demo@example.com"},

		{FirstName: "Sangli", LastName: "Association", Username: "sangli_badminton_association", Password: "password123", Email: "sangli@example.com"},
		{FirstName: "Kolhapur", LastName: "Association", Username: "kolhapur_badminton_association", Password: "password123", Email: "kolhapur@example.com"},
		{FirstName: "Pune", LastName: "Association", Username: "pune_badminton_association", Password: "password123", Email: "pune@example.com"},
		{FirstName: "Mumbai", LastName: "Association", Username: "mumbai_badminton_association", Password: "password123", Email: "mumbai@example.com"},
		{FirstName: "Nagpur", LastName: "Association", Username: "nagpur_badminton_association", Password: "password123", Email: "nagpur@example.com"},
	}
	if err := db.CreateInBatches(organizers, 10); err != nil {
		log.Println(err)
	}
	fmt.Println("org data added successfully")
}
