package entities

import "time"

type Analytics struct {
	URLID      string    `json:"url_id"`
	Timestamp  time.Time `json:"timestamp"`
	Location   string    `json:"location,omitempty"`
	DeviceType string    `json:"device_type,omitempty"`
}