package models

import "time"

// Course - courses in an organization, particular to a program
type Course struct {
	CourseID           string    `gorm:"primary_key;size=50" json:"course_id"`
	InstituteID        string    `gorm:"References:InstituteID;size:50" json:"institute_id"`
	CourseDisciplineID string    `gorm:"References:GraduationProgramID;size:50" json:"course_discipline_id"`
	CourseName         string    `gorm:"size:200" json:"course_name"`
	Duration           int       `gorm:"size=10" json:"duration"`
	Description        string    `gorm:"size:500" json:"description"`
	AddedDate          time.Time `gorm:"size:50" json:"added_date"`
	Status             string    `gorm:"size:20" json:"status"`
	Base
}

// TableName TABLE NAME
func (Course) TableName() string {
	return "base.course"
}
