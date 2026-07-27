package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Ad struct {
	ID uint `gorm:"primaryKey" queryParam:"id"`

	UserID uint `gorm:"not null" queryParam:"user_id"`
	User   User

	ProgramID uint `gorm:"not null" queryParam:"program_id"`
	Program   Program

	Name      string `gorm:"size:100;not null" queryParam:"name"`
	AdType    AdType `gorm:"type:zone_type;not null" queryParam:"ad_type"`
	CreatedAt time.Time `queryParam:"created_at"`
	UpdatedAt time.Time
	Metadata  AdMetaData `gorm:"type:jsonb"`
}

type AdType string

const (
	AdTypeVideo  AdType = "VIDEO"
	AdTypeBanner AdType = "BANNER"
	AdTypeNative AdType = "NATIVE"
)

type AdMetaData struct {
}

func (m *AdMetaData) Scan(value any) error {
	if value == nil {
		*m = AdMetaData{}
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(data, m)
}

func (m *AdMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
