package routers

import (
	"database/sql"
	"encoding/gob"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/music-streaming/internal/config"
	"github.com/tkwang0530/music-streaming/internal/handlers"
	"github.com/tkwang0530/music-streaming/internal/repositories"
	"golang.org/x/oauth2"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	gob.Register(&oauth2.Token{}) // Register oauth2.Token type with gob

	r := gin.Default()
	// Initialize session middleware
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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
	oauth2Group.GET("/google/logout", oauthHandler.Logout)

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
