package repositories

import (
	"database/sql"
	"fmt"

	"github.com/tkwang0530/music-streaming/internal/models"
)

type SongRepository interface {
	Create(song *models.Song) (*models.Song, error)
	GetAll() ([]*models.Song, error)
	GetByID(id int) (*models.Song, error)
	Update(song *models.Song) (*models.Song, error)
	Delete(id int) error
}

type MySQLSongRepository struct {
	db *sql.DB
}

func NewMySQLSongRepository(db *sql.DB) *MySQLSongRepository {
	return &MySQLSongRepository{
		db: db,
	}
}

func (r *MySQLSongRepository) Create(song *models.Song) (*models.Song, error) {
	stmt, err := r.db.Prepare("INSERT INTO songs (title, artist, album, genre, length, year) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, fmt.Errorf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(song.Title, song.Artist, song.Album, song.Genre, song.Length, song.Year)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute SQL statement: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Failed to get last insert ID: %v", err)
	}

	song.ID = int(id)

	return song, nil
}

func (r *MySQLSongRepository) GetAll() ([]*models.Song, error) {
	rows, err := r.db.Query("SELECT * FROM songs")
	if err != nil {
		return nil, fmt.Errorf("Failed to execute SQL query: %v", err)
	}
	defer rows.Close()

	var songs []*models.Song

	for rows.Next() {
		var song models.Song

		if err := rows.Scan(&song.ID, &song.Title, &song.Artist, &song.Album, &song.Genre, &song.Length, &song.Year, &song.CreatedAt, &song.UpdatedAt); err != nil {
			return nil, fmt.Errorf("Failed to scan rows: %v", err)
		}

		songs = append(songs, &song)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Failed to read rows: %v", err)
	}

	return songs, nil
}

func (r *MySQLSongRepository) GetByID(id int) (*models.Song, error) {
	stmt, err := r.db.Prepare("SELECT * FROM songs WHERE id = ?")
	if err != nil {
		return nil, fmt.Errorf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	var song models.Song

	if err := stmt.QueryRow(id).Scan(&song.ID, &song.Title, &song.Artist, &song.Album, &song.Genre, &song.Length, &song.Year, &song.CreatedAt, &song.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, utils.ErrRecordNotFound
		}
		return nil, fmt.Errorf("Failed to execute SQL statement: %v", err)
	}

	return &song, nil
}

func (r *MySQLSongRepository) Update(song *models.Song) (*models.Song, error) {
	stmt, err := r.db.Prepare("UPDATE songs SET title=?, artist=?, album=?, genre=?, length=?, year=? WHERE id=?")
	if err != nil {
		return nil, fmt.Errorf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(song.Title, song.Artist, song.Album, song.Genre, song.Length, song.Year, song.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to execute SQL statement: %v", err)
	}

	return song, nil
}

func (r *MySQLSongRepository) Delete(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM songs WHERE id = ?")
	if err != nil {
		return fmt.Errorf("Failed to prepare SQL statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("Failed to execute SQL statement: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return utils.ErrRecordNotFound
	}

	return nil
}
