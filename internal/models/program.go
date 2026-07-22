package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Program struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null"`
	User   User

	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"type:text"`
	Budget      float64 `gorm:"type:numeric(12,2);not null"`
	BidPrice    float64 `gorm:"type:numeric(10,2);not null"`
	StartDate   *time.Time
	EndDate     *time.Time
	IsActive    bool `gorm:"default:true"`
	IsVerified  bool `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Metadata    ProgramMetaData `gorm:"type:jsonb"`
}

type ProgramMetaData struct {
}

func (m *ProgramMetaData) Scan(value any) error {
	if value == nil {
		*m = ProgramMetaData{}
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(data, m)
}

func (m *ProgramMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
