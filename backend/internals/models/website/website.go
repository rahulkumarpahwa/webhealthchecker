package website

import "time"

type URL struct {
	Domain string `json:"domain"`
	Port   string `json:"port"`
}

type Website struct {
	ID              int64 `json:"id"`
	UserID          int64 `json:"user_id"`
	URL             `json:"url"`
	IntervalSeconds int       `json:"interval_seconds"`
	CreatedAt       time.Time `json:"created_at"`
}
