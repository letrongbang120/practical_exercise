package models

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type Student struct {
	Id        string
	StudentId string
	Name      string
	Email     string
	Pass      string
	CreatedAt time.Time
	UpdatedAt null.Time
	DeletedAt null.Time
}
