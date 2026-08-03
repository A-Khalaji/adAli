package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Ad struct {
	ID uint `gorm:"primaryKey" queryParam:"id" json:"id"`

	UserID uint `gorm:"not null" queryParam:"user_id" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`

	ProgramID uint `gorm:"not null" queryParam:"program_id" json:"program_id"`
	Program   Program `gorm:"foreignKey:ProgramID" json:"program"`

	Name string `gorm:"size:100;not null" queryParam:"name" json:"name"`
	AdType AdType `gorm:"type:ad_type;not null" queryParam:"ad_type" json:"ad_type"`
	CreatedAt time.Time `queryParam:"created_at" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Metadata AdMetaData `gorm:"type:jsonb" json:"metadata"`
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


func (m AdMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}