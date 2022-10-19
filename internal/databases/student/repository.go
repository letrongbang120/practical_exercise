package student

import (
	"context"
	"database/sql"
	"practical_exercise/internal/databases/orm"
	"practical_exercise/internal/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func parseOrmStudent(modelStudent models.Student) orm.Student {
	ormStudent := orm.Student{
		ID:        modelStudent.Id,
		StudentID: null.StringFrom(modelStudent.StudentId),
		Name:      null.StringFrom(modelStudent.Name),
		Email:     null.StringFrom(modelStudent.Email),
		Pass:      null.StringFrom(modelStudent.Pass),
		CreatedAt: modelStudent.CreatedAt,
		UpdatedAt: modelStudent.UpdatedAt,
		DeletedAt: modelStudent.DeletedAt,
	}
	return ormStudent
}

func parseStudent(ormStudent orm.Student) models.Student {
	student := models.Student{
		Id:        ormStudent.ID,
		StudentId: ormStudent.StudentID.String,
		Name:      ormStudent.Name.String,
		Email:     ormStudent.Email.String,
		Pass:      ormStudent.Pass.String,
		CreatedAt: ormStudent.CreatedAt,
		UpdatedAt: ormStudent.UpdatedAt,
		DeletedAt: ormStudent.DeletedAt,
	}
	return student
}

func (management *Management) Create(ctx context.Context, student models.Student) (models.Student, error) {
	ormStudent := parseOrmStudent(student)
	err := ormStudent.Insert(ctx, management.dbconn, boil.Infer())
	if err != nil {
		return models.Student{}, nil
	}
	return parseStudent(ormStudent), nil
}
