package models

import "time"

// Base common fields
type Base struct {
	CreatedBy   string    `gorm:"size=50" json:"created_by"`
	CreatedTime time.Time `gorm:"size=50" json:"created_time"`
	UpdatedBy   string    `gorm:"size=50" json:"updated_by"`
	UpdatedTime time.Time `gorm:"size=50" json:"updated_time"`
}
