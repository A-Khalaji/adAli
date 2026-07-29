package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Site struct {
	ID uint `gorm:"primaryKey" queryParam:"id"`

	UserID uint `gorm:"not null" queryParam:"user_id"`
	User   User

	Name       string `gorm:"size:100;not null" queryParam:"name"`
	Domain     string `gorm:"size:225;not null" queryParam:"domain"`
	Identifier string `gorm:"size:100;not null" queryParam:"identifier"`
	IsActive   bool   `gorm:"default:true" queryParam:"is_active"`
	IsVerified bool   `gorm:"default:true" queryParam:"is_verified"`
	CreatedAt  time.Time `queryParam:"created_at"`
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
