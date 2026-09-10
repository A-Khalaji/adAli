package models

import "time"

type Report struct {
	Date      time.Time `queryParam:"date" json:"date"`
	AdID      uint64    `queryParam:"ad_id" json:"ad_id"`
	ProgramID uint64    `queryParam:"program_id" json:"program_id"`
	ZoneID    uint64    `queryParam:"zone_id" json:"zone_id"`
	SiteID    uint64    `queryParam:"site_id" json:"site_id"`
	Views     uint64    `queryParam:"views" json:"views"`
	Clicks    uint64    `queryParam:"clicks" json:"clicks"`
	ClickRate float64   `json:"click_rate"`
}
