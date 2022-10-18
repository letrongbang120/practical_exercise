package models

import "time"

type Course struct {
	Id           string
	TeacherId    string
	SubjectId    string
	CourseName   string
	StartDay     time.Time
	CourseExtend float32
	CreditAmount int
	SlotAmount   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
