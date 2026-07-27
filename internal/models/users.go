package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID         uint     `gorm:"primaryKey"`
	Name       string   `gorm:"size:100;not null"`
	Email      string   `gorm:"size:255;uniqueIndex;not null"`
	Password   string   `gorm:"size:255;not null"`
	UserType   UserType `gorm:"type:user_Type;not null"`
	IsActive   bool     `gorm:"default:true"`
	IsVerified bool     `gorm:"default:false"`
	CreatedAt  time.Time
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
