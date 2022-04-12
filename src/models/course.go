package models

import "time"

// Course - courses in an organization, particular to a program
type Course struct {
	CourseID           string
	InstituteID        string
	CourseDisciplineID string
	CourseName         string
	Duration           int
	Description        string
	AddedDate          time.Time
	Status             string
	Base
}

// TableName TABLE NAME
func (Course) TableName() string {
	return "base.course"
}
