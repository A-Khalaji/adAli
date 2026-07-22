package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Report struct {
	ID uint `gorm:"primaryKey"`

	ZoneID uint `gorm:"not null"`
	Zone   Zone

	AdID uint `gorm:"not null"`
	Ad   Ad

	Views      uint           `gorm:"default:0"`
	Clicks     uint           `gorm:"default:0"`
	ReportDate time.Time      `gorm:"not null"`
	Metadata   ReportMetaData `gorm:"type:jsonb"`
}

type ReportMetaData struct {
}

func (m *ReportMetaData) Scan(value any) error {
	if value == nil {
		*m = ReportMetaData{}
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(data, m)
}

func (m *ReportMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
