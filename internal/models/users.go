package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID         uint     `gorm:"primaryKey" queryParam:"id"`
	Name       string   `gorm:"size:100;not null" queryParam:"name"`
	Email      string   `gorm:"size:255;uniqueIndex;not null" queryParam:"email"`
	Password   string   `gorm:"size:255;not null"`
	UserType   UserType `gorm:"type:user_Type;not null" queryParam:"user_type"`
	IsActive   bool     `gorm:"default:true" queryParam:"is_active"`
	IsVerified bool     `gorm:"default:false" queryParam:"is_verified"`
	CreatedAt  time.Time `queryParam:"created_at"`
	UpdatedAt  time.Time
	Metadata   UserMetaData `gorm:"type:jsonb"`
}

type UserType string

const (
	RoleAdmin      UserType = "ADMIN"
	RolePublisher  UserType = "PUBLISHER"
	RoleAdvertiser UserType = "ADVERTISER"
)

type UserMetaData struct {
}

func (m *UserMetaData) Scan(value any) error {
	if value == nil {
		*m = UserMetaData{}
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(data, m)
}

func (m *UserMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
