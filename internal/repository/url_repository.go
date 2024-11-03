package repository

import (
	"database/sql"
	"goshort/internal/core/entities"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *entities.URL) error {
	_, err := r.db.Exec(
		"INSERT INTO urls (id, original_url, short_code, created_at, expiry) VALUES (?, ?, ?, ?, ?)",
		url.ID, url.OriginalURL, url.ShortCode, url.CreatedAt, url.Expiry,
	)
	return err
}

func (r *URLRepository) FindByShortCode(shortCode string) (*entities.URL, error) {
	row := r.db.QueryRow("SELECT id, original_url, short_code, created_at, expiry FROM urls WHERE short_code = ?", shortCode)
	url := &entities.URL{}
	if err := row.Scan(&url.ID, &url.OriginalURL, &url.ShortCode, &url.CreatedAt, &url.Expiry); err != nil {
		return nil, err
	}
	return url, nil
}