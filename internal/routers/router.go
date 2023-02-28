package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/tkwang0530/music-streaming/internal/handlers"
)

func NewRouter() *gin.Engine {
	// Create a new Gin router instance
	r := gin.New()

	// Add middleware to the router
	// r.Use(middleware.CORSMiddleware())

	// Create new handlers
	songsHandler := handlers.NewSongsHandler()

	// Add routes to the router
	r.GET("/songs", songsHandler.GetSongs)
	r.POST("/songs", songsHandler.AddSong)
	r.GET("/songs/:id", songsHandler.GetSongByID)
	r.PUT("/songs/:id", songsHandler.UpdateSong)
	r.DELETE("/songs/:id", songsHandler.DeleteSong)

	// Return the router instance
	return r
}
