package repository

import (
	"database/sql"
	"goshort/internal/core/entities"
)

type AnalyticsRepository struct {
	db *sql.DB
}

func NewAnalyticsRepository(db *sql.DB) *AnalyticsRepository {
	return &AnalyticsRepository{db: db}
}

func (r *AnalyticsRepository) LogVisit(analytics *entities.Analytics) error {
	_, err := r.db.Exec(
		"INSERT INTO analytics (url_id, timestamp, location, device_type) VALUES (?, ?, ?, ?)",
		analytics.URLID, analytics.Timestamp, analytics.Location, analytics.DeviceType,
	)
	return err
}

func (r *AnalyticsRepository) GetVisitsByURLID(urlID string) ([]*entities.Analytics, error) {
	rows, err := r.db.Query("SELECT url_id, timestamp, location, device_type FROM analytics WHERE url_id = ?", urlID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var visits []*entities.Analytics
	for rows.Next() {
		analytics := &entities.Analytics{}
		if err := rows.Scan(&analytics.URLID, &analytics.Timestamp, &analytics.Location, &analytics.DeviceType); err != nil {
			return nil, err
		}
		visits = append(visits, analytics)
	}
	return visits, nil
}