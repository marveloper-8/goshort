package domain

import (
	"time"
)

type URL struct {
	ID        string    `json:"id"`
	LongURL   string    `json:"long_url"`
	ShortCode string    `json:"short_code"`
	UserID    string    `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	Clicks    int64     `json:"clicks"`
}

type URLCreateParams struct {
	LongURL   string    `json:"long_url" validate:"required,url"`
	ShortCode string    `json:"short_code,omitempty" validate:"omitempty,alphanum,min=3,max=10"`
	UserID    string    `json:"user_id,omitempty"`
	ExpiresAt time.Time `json:"expires_at,omitempty"`
}
