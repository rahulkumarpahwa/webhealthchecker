package healthcheck

import "time"

type HealthCheck struct {
	ID           int64     `json:"id"`
	WebsiteID    int64     `json:"website_id"`
	Status       bool      `json:"status"`
	ResponseTime int       `json:"response_time"` // milliseconds
	StatusCode   int       `json:"status_code"`
	CheckedAt    time.Time `json:"checked_at"`
}
