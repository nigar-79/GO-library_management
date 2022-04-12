package models

import "time"

// Base common fields
type Base struct {
	CreatedBy   string
	CreatedTime time.Time
	UpdatedBy   string
	UpdatedTime time.Time
}
