package ports

import (
	"context"
	"goshort/internal/core/domain"
	"time"
)

type URLService interface {
	CreateURL(ctx context.Context, params domain.URLCreateParams) (*domain.URL, error)
	GetURL(ctx context.Context, shortcode string) (*domain.URL, error)
	RecordClick(ctx context.Context, shortcode string, clickData domain.ClickEvent) error
}

type AnalyticsService interface {
	GetURLAnalytics(ctx context.Context, urlID string, from, to time.Time) (*domain.AnalyticsSummary, error)
	GetClickEvents(ctx context.Context, urlID string, from, to time.Time) ([]*domain.ClickEvent, error)
}