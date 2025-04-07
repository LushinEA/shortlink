package models

import "time"

type Link struct {
	ID         uint       `json:"links_id"`
	ShortLink  string     `json:"short_link"`
	LongLink   string     `json:"long_link"`
	CreateTS   time.Time  `json:"create_ts"`
	ClickCount int        `json:"click_count"`
	ExpireTS   *time.Time `json:"expire_ts"`
}
