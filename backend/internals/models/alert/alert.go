package alert

import "time"

type Alert struct {
	ID        int64     `json:"id"`
	WebsiteID int64     `json:"website_id"`
	Type      string    `json:"type"`
	SentAt    time.Time `json:"sent_at"`
}
