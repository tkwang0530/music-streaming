package models

import "time"

type Song struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Album     string    `json:"album"`
	Genre     string    `json:"genre"`
	Length    int       `json:"length"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSong(title, artist, album, genre string, length, year int) *Song {
	return &Song{
		Title:  title,
		Artist: artist,
		Album:  album,
		Genre:  genre,
		Length: length,
		Year:   year,
	}
}
