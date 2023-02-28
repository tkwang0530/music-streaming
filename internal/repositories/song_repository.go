package repositories

import (
	"database/sql"

	"github.com/tkwang0530/music-streaming/internal/models"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{
		db: db,
	}
}

func (r *SongRepository) ListSongs() ([]*models.Song, error) {
	rows, err := r.db.Query("SELECT id, title, artist, album, genre, length, year, url, created_at, updated_at FROM songs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	songs := make([]*models.Song, 0)
	for rows.Next() {
		song := &models.Song{}
		if err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Album, &song.Genre, &song.Length, &song.Year, &song.URL, &song.CreatedAt, &song.UpdatedAt); err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}

func (r *SongRepository) GetSong(id string) (*models.Song, error) {
	row := r.db.QueryRow("SELECT id, title, artist, album, genre, length, year, url, created_at, updated_at FROM songs WHERE id=?", id)

	song := &models.Song{}
	err := row.Scan(&song.ID, &song.Title, &song.Artist, &song.Album, &song.Genre, &song.Length, &song.Year, &song.URL, &song.CreatedAt, &song.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (r *SongRepository) CreateSong(params *models.CreateSongParams) error {
	_, err := r.db.Exec("INSERT INTO songs (title, artist, album, genre, length, year, url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())",
		params.Title, params.Artist, params.Album, params.Genre, params.Length, params.Year, params.URL)

	return err
}

func (r *SongRepository) UpdateSong(id int, params *models.CreateSongParams) error {
	_, err := r.db.Exec("UPDATE songs SET title=?, artist=?, album=?, genre=?, length=?, url=?, year=?, updated_at=NOW() WHERE id=?",
		params.Title, params.Artist, params.Album, params.Genre, params.Length, params.URL, params.Year, id)

	return err
}

func (r *SongRepository) DeleteSong(id string) error {
	_, err := r.db.Exec("DELETE FROM songs WHERE id=?", id)

	return err
}
