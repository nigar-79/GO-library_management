package models

import "time"

// GradProgram graduation programs in any organization
type GradProgram struct {
	GraduationProgramID     string    `json:"graduation_program_id"`
	GraduationName          string    `json:"graduation_name"`
	GraduationTypeAddedDate time.Time `json:"graduation_type_added_date"`
	InstituteID             string    `json:"institute_id"`
	Status                  string    `json:"status"`
	Base
}

// TableName TABLE NAME
func (GradProgram) TableName() string {
	return "base.graduation_program"
}
