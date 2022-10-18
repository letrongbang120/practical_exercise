package models

import "time"

type Teacher struct {
	Id        string
	Name      string
	Email     string
	Pass      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
