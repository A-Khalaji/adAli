package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Zone struct {
	ID uint `gorm:"primaryKey" queryParam:"id"`

	UserID uint `gorm:"not null" queryParam:"user_id"`
	User   User

	SiteID uint `gorm:"not null" queryParam:"site_id"`
	Site   Site

	Name       string   `gorm:"size:100;not null" queryParam:"name"`
	ZoneType   ZoneType `gorm:"type:zone_type;not null" queryParam:"zone_type"`
	Identifier string   `gorm:"size:100;not null" queryParam:"identifier"`
	CreatedAt  time.Time `queryParam:"created_at"`
	UpdatedAt  time.Time
	Metadata   ZoneMetaData `gorm:"type:jsonb"`
}

type ZoneType string

const (
	ZoneTypeVideo  ZoneType = "VIDEO"
	ZoneTypeBanner ZoneType = "BANNER"
	ZoneTypeNative ZoneType = "NATIVE"
)

type ZoneMetaData struct {
}

func (m *ZoneMetaData) Scan(value any) error {
	if value == nil {
		*m = ZoneMetaData{}
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(data, m)
}

func (m *ZoneMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
