package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/music-streaming/internal/handlers"
	"github.com/tkwang0530/music-streaming/internal/repositories"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Setup routes
	songsRepo := repositories.NewSongRepository(db)
	songsHandler := handlers.NewSongsHandler(songsRepo)
	r.GET("/songs", songsHandler.ListSongs)
	r.POST("/songs", songsHandler.CreateSong)
	r.GET("/songs/:id", songsHandler.GetSong)
	r.PUT("/songs/:id", songsHandler.UpdateSong)
	r.DELETE("/songs/:id", songsHandler.DeleteSong)

	return r
}
