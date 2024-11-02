package ports

import (
	"context"
	"goshort/internal/core/domain"
	"time"
)

type URLRepository interface {
	Create(ctx context.Context, url *domain.URL) error
	GetByID(ctx context.Context, id string) (*domain.URL, error)
	GetByShortCode(ctx context.Context, shortCode string) (*domain.URL, error)
	Update(ctx context.Context, url *domain.URL) error
	Delete(ctx context.Context, id string) error
	IncrementClicks(ctx context.Context, id string) error
}

type AnalyticsRepository interface {
	SaveClickEvent(ctx context.Context, event *domain.ClickEvent) error
	GetClickEvents(ctx context.Context, urlID string, from, to time.Time) ([]*domain.ClickEvent, error)
	GetAnalyticsSummary(ctx context.Context, urlID string, from, to time.Time) (*domain.AnalyticsSummary, error)
}

type CacheRepository interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
}