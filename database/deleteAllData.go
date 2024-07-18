package database

import (
	"github.com/labstack/echo"
	"net/http"
)

func DeleteData(c echo.Context) error {
	// Start a transaction
	tx := Db.Begin()
	if err := tx.Exec("DELETE FROM matches;").Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if err := tx.Exec("DELETE FROM registrations;").Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if err := tx.Exec("DELETE FROM tournaments;").Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if err := tx.Exec("DELETE FROM players;").Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if err := tx.Exec("DELETE FROM organizers;").Error; err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// Commit the transaction
	tx.Commit()

	return c.JSON(http.StatusOK, echo.Map{"message": "All data deleted successfully"})
}
