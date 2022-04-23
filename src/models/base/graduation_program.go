package models

import (
	"time"

	"sdbgo.in/sdbgo/src/models"
)

// GradProgram graduation programs in any organization
type GradProgram struct {
	GraduationProgramID     string    `gorm:"primary_key;size=50" json:"graduation_program_id"`
	GraduationName          string    `gorm:"size=100" json:"graduation_name"`
	GraduationTypeAddedDate time.Time `gorm:"size=50" json:"graduation_type_added_date"`
	InstituteID             string    `gorm:"References:InstituteID;size:50" json:"institute_id"`
	Status                  string    `gorm:"size=20" json:"status"`
	models.Base
}

// TableName TABLE NAME
func (GradProgram) TableName() string {
	return "base.graduation_program"
}
