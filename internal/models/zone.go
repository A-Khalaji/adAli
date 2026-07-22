package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Zone struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null"`
	User   User

	SiteID uint `gorm:"not null"`
	Site   Site

	Name       string   `gorm:"size:100;not null"`
	ZoneType   ZoneType `gorm:"type:zone_type;not null"`
	Identifier string   `gorm:"size:100;not null"`
	CreatedAt  time.Time
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
