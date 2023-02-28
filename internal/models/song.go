package models

type Song struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
	Genre     string `json:"genre"`
	Length    int    `json:"length"`
	URL       string `json:"url"`
	Year      int    `json:"year"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateSongParams struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	Genre  string `json:"genre"`
	Length int    `json:"length"`
	URL    string `json:"url"`
	Year   int    `json:"year"`
}
