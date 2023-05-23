package handlers

import (
	"go_note/database"
	"go_note/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotesHandler(c *gin.Context) {
	// Fetch all notes from the database
	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var notes []models.Note
	result := db.Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func CreateNoteHandler(c *gin.Context) {
	// Parse request body to get note details
	var requestBody struct {
		Title   string `json:"title"`
		Content string `json:"content"`
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

	newNote := models.Note{
		Title:   requestBody.Title,
		Content: requestBody.Content,
	}
	result := db.Create(&newNote)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note created successfully"})
}

func DeleteNoteHandler(c *gin.Context) {
	// Get the note ID from the URL parameter
	noteID := c.Param("id")

	// Delete the note from the database
	db, err := database.SetupDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	result := db.Delete(&models.Note{}, noteID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}
