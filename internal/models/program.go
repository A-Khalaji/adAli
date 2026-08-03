package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Program struct {
	ID uint `gorm:"primaryKey" queryParam:"id" json:"id"`

	UserID uint `gorm:"not null" queryParam:"user_id" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"user"`

	Name string `gorm:"size:100;not null" queryParam:"name" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	Budget float64 `gorm:"type:numeric(12,2);not null" queryParam:"budget" json:"budget"`
	BidPrice float64 `gorm:"type:numeric(10,2);not null" queryParam:"bid_price" json:"bid_price"`
	StartDate *time.Time `queryParam:"start_date" json:"start_date"`
	EndDate *time.Time `queryParam:"end_date" json:"end_date"`
	IsActive bool `gorm:"default:true" queryParam:"is_active" json:"is_active"`
	IsVerified bool `gorm:"default:true" queryParam:"is_verified" json:"is_verified"`
	CreatedAt time.Time `queryParam:"created_at" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Metadata ProgramMetaData `gorm:"type:jsonb" json:"metadata"`
}


type ProgramMetaData struct {
	Category string `json:"category"`
	Keyword  string `json:"keyword"`
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


func (m ProgramMetaData) Value() (driver.Value, error) {
	return json.Marshal(m)
}