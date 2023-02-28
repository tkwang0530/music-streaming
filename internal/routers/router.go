package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/music-streaming/internal/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Setup middleware
	// r.Use(middleware.AuthMiddleware())

	// Setup routes
	songsHandler := handlers.NewSongsHandler()
	r.GET("/songs", songsHandler.ListSongs)
	r.POST("/songs", songsHandler.CreateSong)
	r.GET("/songs/:id", songsHandler.GetSong)
	r.PUT("/songs/:id", songsHandler.UpdateSong)
	r.DELETE("/songs/:id", songsHandler.DeleteSong)

	return r
}
