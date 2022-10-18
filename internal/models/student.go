package models

import "time"

type Student struct {
	Id        string
	StudentId string
	Name      string
	Email     string
	Pass      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
