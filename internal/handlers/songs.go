package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/music-streaming/internal/models"
	"github.com/tkwang0530/music-streaming/internal/repositories"
	"github.com/tkwang0530/music-streaming/internal/utils"
)

type SongHandler struct {
	songRepository repositories.SongRepository
}

func NewSongHandler(songRepository repositories.SongRepository) *SongHandler {
	return &SongHandler{
		songRepository: songRepository,
	}
}

func (h *SongHandler) GetSongs(c *gin.Context) {
	songs, err := h.songRepository.GetAll()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *SongHandler) AddSong(c *gin.Context) {
	var song models.Song
	if err := json.NewDecoder(c.Request.Body).Decode(&song); err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	if err := h.songRepository.Create(&song); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, song)
}

func (h *SongHandler) GetSongByID(c *gin.Context) {
	id := c.Param("id")

	song, err := h.songRepository.GetByID(id)
	if err != nil {
		utils.HandleError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, song)
}
