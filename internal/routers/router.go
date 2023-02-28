package routers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/music-streaming/internal/config"
	"github.com/tkwang0530/music-streaming/internal/handlers"
	"github.com/tkwang0530/music-streaming/internal/repositories"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// OAuth2 routes
	oauthHandler := handlers.NewAuthHandler(cfg)
	oauth2Group := r.Group("/oauth2")
	oauth2Group.GET("/google/login", oauthHandler.Login)
	oauth2Group.GET("/google/callback", oauthHandler.Callback)

	// Setup routes
	songsRepo := repositories.NewSongRepository(db)
	songsHandler := handlers.NewSongsHandler(songsRepo)
	songsGroup := r.Group("/songs")
	songsGroup.GET("", songsHandler.ListSongs)
	songsGroup.POST("", songsHandler.CreateSong)
	songsGroup.GET("/:id", songsHandler.GetSong)
	songsGroup.PUT("/:id", songsHandler.UpdateSong)
	songsGroup.DELETE("/:id", songsHandler.DeleteSong)

	return r
}
