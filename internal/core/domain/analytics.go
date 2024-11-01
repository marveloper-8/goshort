package domain

import (
	"time"
)

type ClickEvent struct {
	ID         string                 `json:"id"`
	URLId      string                 `json:"url_id"`
	Timestamp  time.Time              `json:"timestamp"`
	IP         string                 `json:"ip"`
	UserAgent  string                 `json:"user_agent"`
	Country    string                 `json:"country"`
	City       string                 `json:"city"`
	Device     string                 `json:"device"`
	Browser    string                 `json:"browser"`
	Platform   string                 `json:"platform"`
	Referrer   string                 `json:"referrer"`
	CustomData map[string]interface{} `json:"custom_data,omitempty"`
}

type AnalyticsSummary struct {
	TotalClicks    int64            `json:"total_clicks"`
	UniqueVisitors int64            `json:"unique_visitors"`
	TopCountries   []CountryStat    `json:"top_countries"`
	TopDevices     []DeviceStat     `json:"top_devices"`
	ClicksOverTime []TimeseriesStat `json:"clicks_over_time"`
}

type CountryStat struct {
	Country string `json:"country"`
	Count   int64  `json:"count"`
}

type DeviceStat struct {
	Device string `json:"device"`
	Count  int64  `json:"count"`
}

type TimeseriesStat struct {
	Timestamp time.Time `json:"timestamp"`
	Count     int64     `json:"count"`
}
