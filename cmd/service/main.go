package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Song struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Length int    `json:"length"`
	Year   int    `json:"year"`
}

var songs []*Song = []*Song{}

func main() {
	r := gin.Default()

	r.GET("/songs", getSongs)
	r.POST("/songs", addSong)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}

func getSongs(c *gin.Context) {
	c.JSON(http.StatusOK, songs)
}

func addSong(c *gin.Context) {
	song := Song{}
	if err := json.NewDecoder(c.Request.Body).Decode(&song); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	songs = append(songs, &song)
	c.JSON(http.StatusOK, song)
}
