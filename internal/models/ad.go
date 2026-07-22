package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Ad struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null"`
	User   User

	ProgramID uint `gorm:"not null"`
	Program   Program

	Name      string `gorm:"size:100;not null"`
	AdType    AdType `gorm:"type:zone_type;not null"`
	CreatedAt time.Time
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
