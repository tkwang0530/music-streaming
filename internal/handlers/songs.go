package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/music-streaming/internal/models"
	"github.com/tkwang0530/music-streaming/internal/repositories"
)

type SongsHandler struct {
	songsRepo *repositories.SongRepository
}

func NewSongsHandler(songsRepo *repositories.SongRepository) *SongsHandler {
	return &SongsHandler{
		songsRepo: songsRepo,
	}
}

func (h *SongsHandler) ListSongs(c *gin.Context) {
	songs, err := h.songsRepo.ListSongs()
	fmt.Println("err: ", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *SongsHandler) GetSong(c *gin.Context) {
	id := c.Param("id")

	song, err := h.songsRepo.GetSong(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get song"})
		return
	}

	if song == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *SongsHandler) CreateSong(c *gin.Context) {
	params := &models.CreateSongParams{}
	if err := c.BindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.songsRepo.CreateSong(params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create song"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *SongsHandler) UpdateSong(c *gin.Context) {
	id := c.Param("id")

	song, err := h.songsRepo.GetSong(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get song"})
		return
	}

	if song == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Song not found"})
		return
	}

	params := &models.CreateSongParams{}
	if err := c.BindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.songsRepo.UpdateSong(song.ID, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
		return
	}

	c.JSON(http.StatusOK, song)
}

func (h *SongsHandler) DeleteSong(c *gin.Context) {
	id := c.Param("id")

	if err := h.songsRepo.DeleteSong(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted"})
}
