package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey" queryParam:"id" json:"id"`
	Name string `gorm:"size:100;not null" queryParam:"name" json:"name"`
	Email string `gorm:"size:255;uniqueIndex;not null" queryParam:"email" json:"email"`
	Password string `gorm:"size:255;not null" json:"password"`
	UserType UserType `gorm:"type:user_type;not null" queryParam:"user_type" json:"user_type"`
	IsActive bool `gorm:"default:true" queryParam:"is_active" json:"is_active"`
	IsVerified bool `gorm:"default:false" queryParam:"is_verified" json:"is_verified"`
	CreatedAt time.Time `queryParam:"created_at" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Metadata UserMetaData `gorm:"type:jsonb" json:"metadata"`
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

func (m UserMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}
