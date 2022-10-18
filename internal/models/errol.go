package models

import "time"

type Errol struct {
	Id        string
	StudentId string
	CourseId  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
