package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Site struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null"`
	User   User

	Name       string `gorm:"size:100;not null"`
	Domain     string `gorm:"size:225;not null"`
	Identifier string `gorm:"size:100;not null"`
	IsActive   bool   `gorm:"default:true"`
	IsVerified bool   `gorm:"default:true"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Metadata   SiteMetaData `gorm:"type:jsonb"`
}

type SiteMetaData struct {
}

func (m *SiteMetaData) Scan(value any) error {
	if value == nil {
		*m = SiteMetaData{}
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(data, m)
}

func (m *SiteMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
