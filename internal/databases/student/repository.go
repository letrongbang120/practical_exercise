package student

import (
	"context"
	"database/sql"
	"practical_exercise/internal/models"
)

type Repository interface {
	Create(ctx context.Context, student models.Student) (models.Student, error)
}

type Management struct {
	dbconn *sql.DB
}

func NewManagement(dbconn *sql.DB) Repository {
	return &Management{
		dbconn: dbconn,
	}
}

func (management *Management) Create(ctx context.Context, student models.Student) (models.Student, error) {
	return student, nil
}
