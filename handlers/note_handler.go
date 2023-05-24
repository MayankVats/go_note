package handlers

import (
	"go_note/database"
	"go_note/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetNotesHandler(c *gin.Context) {
	var requestBody struct {
		Sid string `json:"sid"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Body"})
		return
	}

	// Fetch all notes from the database
	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var session models.Session
	sessionResult := db.Where("id = ?", requestBody.Sid).First(&session)
	if sessionResult.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	if session.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	type Note struct {
		ID   uint32
		Note string
	}
	var notes []Note

	result := db.Where("user_id = ?", session.UserID).Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func CreateNoteHandler(c *gin.Context) {
	// Parse request body to get note details
	var requestBody struct {
		Sid  string `json:"sid"`
		Note string `json:"note"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Create a new note in the database
	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var session models.Session
	sessionResult := db.Where("id = ?", requestBody.Sid).First(&session)
	if sessionResult.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	if session.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	newNote := models.Note{
		UserID: session.UserID,
		Note:   requestBody.Note,
	}
	result := db.Create(&newNote)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": newNote.ID})
}

func DeleteNoteHandler(c *gin.Context) {
	var requestBody struct {
		Sid    string `json:"sid"`
		NoteID uint32 `json:"id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Delete the note from the database
	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var session models.Session
	sessionResult := db.Where("id = ?", requestBody.Sid).First(&session)
	if sessionResult.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	if session.ExpiresAt.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authorized"})
		return
	}

	result := db.Unscoped().Delete(&models.Note{}, requestBody.NoteID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
