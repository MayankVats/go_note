package main

import (
	"go_note/database"
	"go_note/handlers"
	"go_note/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db, err := database.SetupDatabase()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Auto-migrate the models
	db.AutoMigrate(&models.User{}, &models.Note{})

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "notes app backend-service")
	})

	// User handlers
	router.POST("/signup", handlers.SignupHandler)
	router.POST("/login", handlers.LoginHandler)

	// Note handlers
	router.GET("/notes", handlers.GetNotesHandler)
	router.POST("/notes", handlers.CreateNoteHandler)
	router.DELETE("/notes/:id", handlers.DeleteNoteHandler)

	router.Run(":3000")
}
