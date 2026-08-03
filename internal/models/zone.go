package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Zone struct {
	ID uint `gorm:"primaryKey" queryParam:"id" json:"id"`

	UserID uint `gorm:"not null" queryParam:"user_id" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`

	SiteID uint `gorm:"not null" queryParam:"site_id" json:"site_id"`
	Site   Site `gorm:"foreignKey:SiteID" json:"site"`

	Name       string   `gorm:"size:100;not null" queryParam:"name" json:"name"`
	ZoneType   ZoneType `gorm:"type:zone_type;not null" queryParam:"zone_type" json:"zone_type"`
	Identifier string   `gorm:"size:100;not null" queryParam:"identifier" json:"identifier"`

	CreatedAt time.Time `queryParam:"created_at" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Metadata ZoneMetaData `gorm:"type:jsonb" json:"metadata"`
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
